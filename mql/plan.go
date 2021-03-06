// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package mql

// Plan is the interface of SQL execution
type Plan interface {
	execute()
}
