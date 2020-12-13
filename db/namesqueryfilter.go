// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// NamesQueryFilter ...
type NamesQueryFilter struct {
	key     string
	path    *QueryPath
	columns [][]byte
}

// NewNamesQueryFilter ...
func NewNamesQueryFilter(key string, columnParent *QueryPath, column []byte) QueryFilter {
	n := &NamesQueryFilter{}
	n.key = key
	n.path = columnParent
	n.columns = make([][]byte, 1)
	n.columns = append(n.columns, column)
	return n
}

// NewNamesQueryFilterS ...
func NewNamesQueryFilterS(key string, columnParent *QueryPath, columns [][]byte) QueryFilter {
	n := &NamesQueryFilter{}
	n.key = key
	n.path = columnParent
	n.columns = columns
	return n
}
func (n *NamesQueryFilter) getKey() string {
	return n.key
}
func (n *NamesQueryFilter) getPath() *QueryPath {
	return n.path
}

func containsC(list [][]byte, name string) bool {
	for _, elem := range list {
		if string(elem) == name {
			return true
		}
	}
	return false
}
func (n *NamesQueryFilter) filterSuperColumn(superColumn SuperColumn, gcBefore int) SuperColumn {
	for name := range superColumn.getSubColumns() {
		if containsC(n.columns, name) == false {
			superColumn.Remove(name)
		}
	}
	return superColumn
}

func (n *NamesQueryFilter) getMemColumnIterator(memtable *Memtable) ColumnIterator {
	return memtable.getNamesIterator(n)
}

func (n *NamesQueryFilter) getSSTableColumnIterator(sstable *SSTableReader) ColumnIterator {
	return NewSSTableNamesIterator(sstable, n.key, n.columns)
}

func (n *NamesQueryFilter) collectCollatedColumns(returnCF *ColumnFamily, collatedColumns *CollatedIterator, gcBefore int) {
	return
}
