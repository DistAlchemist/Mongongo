// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import "math"

// IdentityQueryFilter ...
type IdentityQueryFilter struct {
	*SliceQueryFilter
	key  string
	path *QueryPath
}

// NewIdentityQueryFilter ...
func NewIdentityQueryFilter(key string, path *QueryPath) QueryFilter {
	p := &IdentityQueryFilter{}
	p.SliceQueryFilter = NewSliceQueryFilter(key, path, nil, nil, false, math.MaxInt32)
	return p
}

func (p *IdentityQueryFilter) filterSuperColumn(superColumn SuperColumn, gcBefore int) SuperColumn {
	// no filtering done, deliberately
	return superColumn
}
