// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

import "github.com/DistAlchemist/Mongongo/db"

// ColumnOrSuperColumn ...
type ColumnOrSuperColumn struct {
	Column  *db.Column
	SColumn *db.SuperColumn
}

// NewColumnOrSuperColumn ...
func NewColumnOrSuperColumn(column *db.Column, superColumn *db.SuperColumn) ColumnOrSuperColumn {
	res := ColumnOrSuperColumn{}
	res.Column = column
	res.SColumn = superColumn
	return res
}
