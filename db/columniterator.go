// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// ColumnIterator ...
type ColumnIterator interface {
	getColumnFamily() *ColumnFamily
	close()
	hasNext() bool
	next() IColumn
}

// SimpleColumnIterator ...
type SimpleColumnIterator struct {
	curIndex     int
	columnFamily *ColumnFamily
	columns      [][]byte
}

// NewSColumnIterator ...
func NewSColumnIterator(curIndex int, cf *ColumnFamily, columns [][]byte) *SimpleColumnIterator {
	c := &SimpleColumnIterator{}
	c.curIndex = 0
	c.columnFamily = cf
	c.columns = columns
	return c
}

func (c *SimpleColumnIterator) getColumnFamily() *ColumnFamily {
	return c.columnFamily
}
func (c *SimpleColumnIterator) close() {
	return
}
func (c *SimpleColumnIterator) hasNext() bool {
	return c.curIndex < len(c.columns)
}
func (c *SimpleColumnIterator) next() IColumn {
	if c.hasNext() == false {
		return nil
	}
	c.curIndex++
	return c.columnFamily.GetColumn(string(c.columns[c.curIndex]))
}

// AbstractColumnIterator ...
type AbstractColumnIterator struct {
	curIndex     int
	columnFamily *ColumnFamily
	columns      []IColumn
}

// NewAColumnIterator ...
func NewAColumnIterator(curIndex int, columnFamily *ColumnFamily,
	columns []IColumn) *AbstractColumnIterator {
	c := &AbstractColumnIterator{}
	c.curIndex = curIndex
	c.columnFamily = columnFamily
	c.columns = columns
	return c
}

func (c *AbstractColumnIterator) getColumnFamily() *ColumnFamily {
	return c.columnFamily
}

func (c *AbstractColumnIterator) close() {}

func (c *AbstractColumnIterator) hasNext() bool {
	return c.curIndex < len(c.columns)
}

func (c *AbstractColumnIterator) next() IColumn {
	if c.hasNext() == false {
		return nil
	}
	c.curIndex++
	return c.columns[c.curIndex-1]
}
