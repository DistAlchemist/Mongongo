// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// IColumn provide interface for Column and SuperColumn
type IColumn interface {
	addColumn(name string, column IColumn)
	getSize() int32
}
