// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package dht

// OrderPreservingPartitioner ...
type OrderPreservingPartitioner struct{}

// NewOPP ...
func NewOPP() *OrderPreservingPartitioner {
	r := &OrderPreservingPartitioner{}
	return r
}

// DecorateKey ...
func (r *OrderPreservingPartitioner) DecorateKey(key string) string {
	return key
}

// UndecorateKey ...
func (r *OrderPreservingPartitioner) UndecorateKey(decoratedKey string) string {
	return decoratedKey
}

// Compare ...
func (r *OrderPreservingPartitioner) Compare(s1, s2 string) int {
	if s1 == s2 {
		return 0
	} else if s1 < s2 {
		return -1
	}
	return 1
}
