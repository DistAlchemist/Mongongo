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
	return ""
}
func (n *NamesQueryFilter) getPath() *QueryPath {
	return nil
}
func (n *NamesQueryFilter) filterSuperColumn(superColumn SuperColumn, gcBefore int) SuperColumn {
	return SuperColumn{}
}
func (n *NamesQueryFilter) getMemColumnIterator(memtable *Memtable) ColumnIterator {
	return nil
}
func (n *NamesQueryFilter) getSSTableColumnIterator(sstable *SSTableReader) ColumnIterator {
	return nil
}
func (n *NamesQueryFilter) collectCollatedColumns(returnCF *ColumnFamily, collatedColumns *CollatedIterator, gcBefore int) {
	return
}
