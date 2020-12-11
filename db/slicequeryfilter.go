// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import "log"

// SliceQueryFilter ...
type SliceQueryFilter struct {
	*AQueryFilter
	start    []byte
	finish   []byte
	reversed bool
	count    int
}

// NewSliceQueryFilter ...
func NewSliceQueryFilter(key string, columnParent *QueryPath, start, finish []byte,
	reversed bool, count int) *SliceQueryFilter {
	s := &SliceQueryFilter{}
	s.AQueryFilter = NewAQueryFilter(key, columnParent)
	s.start = start
	s.finish = finish
	s.reversed = reversed
	s.count = count
	return s
}

func (s *SliceQueryFilter) getMemColumnIterator(memtable *Memtable) ColumnIterator {
	return memtable.getSliceIterator(s)
}

func (s *SliceQueryFilter) getSSTableColumnIterator(sstable *SSTableReader) ColumnIterator {
	return NewSSTableSliceIterator(sstable, s.key, s.start, s.reversed)
}

func (s *SliceQueryFilter) collectCollatedColumns(returnCF *ColumnFamily, collatedColumns *CollatedIterator, gcBefore int) {
	// define a 'reduced' iterator that merges columns with the same
	// name, which greatly simplies computing liveColumns in the
	// presence of tombstones.
	// BUT I will omit this part :)
	// TODO make a reduce iterator
	s.collectReducedColumns(returnCF, collatedColumns, gcBefore)
}

func (s *SliceQueryFilter) collectReducedColumns(returnCF *ColumnFamily, reducedColumns *CollatedIterator, gcBefore int) {
	liveColumns := 0
	for {
		column := reducedColumns.next()
		if column == nil {
			break
		}
		if liveColumns >= s.count {
			break
		}
		log.Printf("collecting columns\n")
		if len(s.finish) > 0 &&
			(!s.reversed && column.getName() > string(s.finish) ||
				s.reversed && column.getName() < string(s.finish)) {
			break
		}
		// only count live columns towards the `count` criteria
		if !column.isMarkedForDelete() &&
			(!returnCF.isMarkedForDelete() || column.mostRecentChangeAt() > returnCF.getMarkedForDeleteAt()) {
			liveColumns++
		}
		// but we need to add all non-gc-able columns to the result of read repair:
		// the column itself must be not gc-able, (1)
		// and if its container is deleted, the column must be changed more
		// recently than the container tombstone (2)
		// (since otherwise, the only thing repair cares about is the container tombstone)
		if (!column.isMarkedForDelete() || column.getLocalDeletionTime() > gcBefore) && // (1)
			(!returnCF.isMarkedForDelete() || column.mostRecentChangeAt() > returnCF.markedForDeleteAt) { // (2)
			returnCF.addColumn(column)
		}
	}
}
