// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

import (
	"log"

	"github.com/DistAlchemist/Mongongo/db"
)

// Mongongo expose the interface of operations
type Mongongo struct {
	// Mongongo struct
	Hostname string
	Port     int
	// storageService *StorageService
}

// // ExecuteArgs arguments of executeQueryOnServer
// type ExecuteArgs struct {
// 	Line string
// }

// // ExecuteReply reply format of executeQueryOnServer
// type ExecuteReply struct {
// 	Result mql.Result
// }

// // ExecuteQueryOnServer handles the rpc from cli client
// func (mg *Mongongo) ExecuteQueryOnServer(args *ExecuteArgs, reply *ExecuteReply) error {
// 	//
// 	line := args.Line
// 	log.Printf("server executing %+v\n", line)
// 	reply.Result = mql.ExecuteQuery(line)
// 	return nil
// }

// Start setup other service such as storageService
func (mg *Mongongo) Start() {
	GetInstance().Start()
}

// InsertArgs ...
type InsertArgs struct {
	Table            string
	Key              string
	CPath            ColumnPath
	Value            []byte
	Timestamp        int64
	ConsistencyLevel int
}

// InsertReply ...
type InsertReply struct {
	Result string
}

// Insert is an rpc
func (mg *Mongongo) Insert(args *InsertArgs, reply *InsertReply) error {
	log.Printf("enter mg.Insert\n")
	table := args.Table
	key := args.Key
	columnPath := args.CPath
	value := args.Value
	timestamp := args.Timestamp
	consistencyLevel := args.ConsistencyLevel
	rm := db.NewRowMutation(table, key)
	rm.AddQ(db.NewQueryPath(columnPath.ColumnFamily, columnPath.SuperColumn, columnPath.Column),
		value, timestamp)
	mg.doInsert(consistencyLevel, rm)
	return nil
}

func (mg *Mongongo) doInsert(consistencyLevel int, rm db.RowMutation) {
	if consistencyLevel != 0 {
		insertBlocking(rm, consistencyLevel)
	} else {
		insert(rm)
	}
}

// GetSliceArgs ...
type GetSliceArgs struct {
	Keyspace         string
	Key              string
	ColumnParent     ColumnParent
	Predicate        SlicePredicate
	ConsistencyLevel int
}

// GetSliceReply ...
type GetSliceReply struct {
	Columns []ColumnOrSuperColumn
}

// GetSlice ...
func (mg *Mongongo) GetSlice(args *GetSliceArgs, reply *GetSliceReply) error {
	log.Printf("enter mg.GetSlice\n")
	keyspace := args.Keyspace
	key := args.Key
	columnParent := args.ColumnParent
	predicate := args.Predicate
	consistencyLevel := args.ConsistencyLevel
	reply.Columns = mg.multigetSliceInternal(keyspace, []string{key}, columnParent,
		predicate, consistencyLevel)[key]
	return nil
}

func (mg *Mongongo) multigetSliceInternal(keyspace string, keys []string, columnParent ColumnParent,
	predicate SlicePredicate, consistencyLevel int) map[string][]ColumnOrSuperColumn {
	commands := make([]db.ReadCommand, 0)
	sRange := predicate.SRange
	if predicate.ColumnNames != nil {
		for _, key := range keys {
			commands = append(commands, db.NewSliceByNamesReadCommand(keyspace, key,
				*db.NewQueryPath(columnParent.ColumnFamily, columnParent.SuperColumn, nil), predicate.ColumnNames))
		}
	} else {
		for _, key := range keys {
			commands = append(commands, db.NewSliceFromReadCommand(keyspace, key,
				*db.NewQueryPath(columnParent.ColumnFamily, columnParent.SuperColumn, nil),
				sRange.Start, sRange.Finish, sRange.Reversed, sRange.Count))
		}
	}
	return mg.getSlice(commands, consistencyLevel)
}

func (mg *Mongongo) getSlice(commands []db.ReadCommand, consistencyLevel int) map[string][]ColumnOrSuperColumn {
	cfs := mg.readColumnFamily(commands, consistencyLevel)
	cfMap := make(map[string][]ColumnOrSuperColumn)
	for _, command := range commands {
		cf := cfs[command.GetKey()]
		_, ok := command.(*db.SliceFromReadCommand)
		reverseOrder := false
		if ok && command.(*db.SliceFromReadCommand).Reversed {
			reverseOrder = true
		}
		if cf == nil || len(cf.Columns) == 0 {
			cfMap[command.GetKey()] = nil
			continue
		}
		if command.GetQPath().SuperColumnName != nil {
			var column db.IColumn
			for _, column = range cf.Columns {
				break
			}
			subColumns := column.GetSubColumns()
			if subColumns == nil || len(subColumns) == 0 {
				cfMap[command.GetKey()] = nil
				continue
			}
			cl := make([]db.IColumn, 0)
			for _, c := range subColumns {
				cl = append(cl, c)
			}
			cfMap[command.GetKey()] = mg.procColumns(cl, reverseOrder)
			continue
		}
		if cf.IsSuper() {
			cfMap[command.GetKey()] = mg.procSuperColumns(cf.GetSortedColumns(), reverseOrder)
		} else {
			cfMap[command.GetKey()] = mg.procColumns(cf.GetSortedColumns(), reverseOrder)
		}
	}
	return cfMap
}

func (mg *Mongongo) procColumns(columns []db.IColumn, reverseOrder bool) []ColumnOrSuperColumn {
	res := make([]ColumnOrSuperColumn, len(columns))
	for _, column := range columns {
		if column.IsMarkedForDelete() {
			continue
		}
		c := db.NewColumn(column.GetName(), string(column.GetValue()), column.GetTimestamp(), false)
		res = append(res, NewColumnOrSuperColumn(&c, nil))
	}
	if reverseOrder {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	return res
}

func (mg *Mongongo) procSuperColumns(columns []db.IColumn, reverseOrder bool) []ColumnOrSuperColumn {
	res := make([]ColumnOrSuperColumn, len(columns))
	for _, column := range columns {
		subcolumns := mg.procSubColumns(column.GetSubColumns())
		if len(subcolumns) == 0 {
			continue
		}
		c := db.NewSuperColumnN(column.GetName(), subcolumns)
		res = append(res, NewColumnOrSuperColumn(nil, &c))
	}
	if reverseOrder {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	return res
}

func (mg *Mongongo) procSubColumns(columns map[string]db.IColumn) map[string]db.IColumn {
	if columns == nil || len(columns) == 0 {
		return nil
	}
	res := make(map[string]db.IColumn)
	for name, column := range columns {
		if column.IsMarkedForDelete() {
			continue
		}
		c := db.NewColumn(name, string(column.GetValue()), column.GetTimestamp(), false)
		res[name] = c
	}
	return res
}

func (mg *Mongongo) readColumnFamily(commands []db.ReadCommand, consistencyLevel int) map[string]*db.ColumnFamily {
	cfName := commands[0].GetCFName()
	res := make(map[string]*db.ColumnFamily)
	if consistencyLevel == 0 {
		log.Fatal("consistency level 0 may not be applied to read operation")
	}
	if consistencyLevel == 3 {
		log.Fatal("consistency level all is not yet supported on read operation")
	}
	rows := readProtocol(commands, consistencyLevel)
	for _, row := range rows {
		res[row.Key] = row.ColumnFamilies[cfName]
	}
	return res
}

// GetArgs ...
type GetArgs struct {
	Keyspace         string
	Key              string
	ColumnPath       ColumnPath
	ConsistencyLevel int
}

// GetReply ...
type GetReply struct {
	Cosc ColumnOrSuperColumn
}

// Get ...
func (mg *Mongongo) Get(args *GetArgs, reply *GetReply) error {
	log.Printf("enter mg.Get\n")
	keyspace := args.Keyspace
	key := args.Key
	columnPath := args.ColumnPath
	consistencyLevel := args.ConsistencyLevel
	reply.Cosc = mg.multigeteInternal(keyspace, []string{key}, columnPath,
		consistencyLevel)[key]
	return nil
}

func (mg *Mongongo) multigeteInternal(table string, keys []string, columnPath ColumnPath,
	consistencyLevel int) map[string]ColumnOrSuperColumn {
	path := db.NewQueryPath(columnPath.ColumnFamily, []byte(columnPath.SuperColumn),
		[]byte(columnPath.Column))
	// assume without super column, just
	// get table.cf['key']['column']
	name := columnPath.Column
	commands := make([]db.ReadCommand, 0)
	for _, key := range keys {
		commands = append(commands, db.NewSliceByNamesReadCommand(table, key, *path, [][]byte{name}))
	}
	cfMap := make(map[string]ColumnOrSuperColumn)
	columnsMap := mg.multigetColumns(commands, consistencyLevel)
	for _, command := range commands {
		columns := columnsMap[command.GetKey()]
		var c ColumnOrSuperColumn
		if columns == nil {
			c = ColumnOrSuperColumn{}
		} else {
			var column db.IColumn
			for _, column = range columns {
				break
			}
			if column.IsMarkedForDelete() {
				c = ColumnOrSuperColumn{}
			} else {
				nc := db.NewColumn(column.GetName(), string(column.GetValue()),
					column.GetTimestamp(), false)
				c = NewColumnOrSuperColumn(&nc, nil)
			}
		}
		cfMap[command.GetKey()] = c
	}
	return cfMap
}

func (mg *Mongongo) multigetColumns(commands []db.ReadCommand, consistencyLevel int) map[string][]db.IColumn {
	cfs := mg.readColumnFamily(commands, consistencyLevel)
	cfMap := make(map[string][]db.IColumn)
	for _, command := range commands {
		cf := cfs[command.GetKey()]
		if cf == nil {
			continue
		}
		columns := make([]db.IColumn, 0)
		if command.GetQPath().SuperColumnName != nil {
			column := cf.GetColumn(string(command.GetQPath().SuperColumnName))
			if column != nil {
				cms := column.GetSubColumns()
				for _, c := range cms {
					columns = append(columns, c)
				}
			}
		} else {
			columns = cf.GetSortedColumns()
		}
		if columns != nil && len(columns) != 0 {
			cfMap[command.GetKey()] = columns
		}
	}
	return cfMap
}
