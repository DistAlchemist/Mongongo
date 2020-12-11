// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// QueryFilter ...
type QueryFilter interface {
	getKey() string
	getPath() *QueryPath
	filterSuperColumn(superColumn SuperColumn, gcBefore int) SuperColumn
	getMemColumnIterator(memtable *Memtable) ColumnIterator
	getSSTableColumnIterator(sstable *SSTableReader) ColumnIterator
	collectCollatedColumns(returnCF *ColumnFamily, collatedColumns *CollatedIterator, gcBefore int)
}

// AQueryFilter ...
type AQueryFilter struct {
	key  string
	path *QueryPath
}

// NewAQueryFilter ...
func NewAQueryFilter(key string, path *QueryPath) *AQueryFilter {
	p := &AQueryFilter{}
	p.key = key
	p.path = path
	return p
}

func (p *AQueryFilter) getKey() string {
	return p.key
}

func (p *AQueryFilter) getPath() *QueryPath {
	return p.path
}

func (p *AQueryFilter) filterSuperColumn(superColumn SuperColumn, gcBefore int) SuperColumn {
	return superColumn
}

func (p *AQueryFilter) getMemColumnIterator(memtable *Memtable) ColumnIterator {
	return nil
}

func (p *AQueryFilter) getSSTableColumnIterator(sstable *SSTableReader) ColumnIterator {
	return nil
}
