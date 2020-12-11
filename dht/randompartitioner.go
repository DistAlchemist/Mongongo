// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package dht

import (
	"crypto/md5"
	"strings"
)

// RandomPartInstance ...
var RandomPartInstance = NewRandomPartitioner()

// RandomPartitioner ...
type RandomPartitioner struct {
}

// NewRandomPartitioner ...
func NewRandomPartitioner() *RandomPartitioner {
	r := &RandomPartitioner{}
	return r
}

func hash(key string) string {
	tmp := md5.Sum([]byte(key))
	return string(tmp[:])
}

// DecorateKey ...
func (r *RandomPartitioner) DecorateKey(key string) string {
	return hash(key) + ":" + key
}

// UndecorateKey ...
func (r *RandomPartitioner) UndecorateKey(decoratedKey string) string {
	parts := strings.Split(decoratedKey, ":")
	return parts[1]
}

// Compare ...
func (r *RandomPartitioner) Compare(s1, s2 string) int {
	if s1 == s2 {
		return 0
	} else if s1 < s2 {
		return -1
	}
	return 1
}
