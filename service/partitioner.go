// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

type iPartitioner interface {
	hash(key string) uint64
}

type orderPreservingHashPartitioner struct{}

func (p orderPreservingHashPartitioner) hash(key string) uint64 {
	return 0
}

type randomPartitioner struct{}

func (p randomPartitioner) hash(key string) uint64 {
	return 0
}
