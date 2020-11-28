// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// SuperColumn implements IColumn interface
type SuperColumn struct {
	Name    string
	Columns map[string]IColumn
}

func (sc SuperColumn) addColumn(name string, column IColumn) {
	if sc.Columns == nil {
		sc.Columns = make(map[string]IColumn)
	}
	sc.Columns[name] = column
}

// NewSuperColumn constructs a SuperColun
func NewSuperColumn(name string) SuperColumn {
	sc := SuperColumn{}
	sc.Name = name
	return sc
}
