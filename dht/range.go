// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package dht

// Range is a representation of the range that
// a node is responsible for on the DHT ring.
type Range struct {
	left  string
	right string
}
