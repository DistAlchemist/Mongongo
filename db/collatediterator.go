// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// CollatedIterator ...
type CollatedIterator struct {
	iterators []ColumnIterator
	buf       map[int]IColumn
}

// NewCollatedIterator ...
func NewCollatedIterator(iterators []ColumnIterator) *CollatedIterator {
	c := &CollatedIterator{}
	c.iterators = iterators
	c.buf = make(map[int]IColumn)
	for i := range iterators {
		c.buf[i] = nil
	}
	return c
}

func (c *CollatedIterator) next() IColumn {
	minIndex := 0
	var minCol IColumn
	nilCnt := 0
	for i := range c.buf {
		if c.buf[i] == nil {
			c.buf[i] = c.iterators[i].next()
		}
		if c.buf[i] == nil {
			nilCnt++
		} else {
			if minCol == nil {
				minIndex = i
				minCol = c.buf[i]
			} else {
				if c.buf[i].getName() < minCol.getName() {
					minIndex = i
					minCol = c.buf[i]
				}
			}
		}
	}
	c.buf[minIndex] = nil
	return minCol
}
