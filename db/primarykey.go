// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"crypto/md5"
	"sort"

	"github.com/DistAlchemist/Mongongo/config"
)

// PrimaryKey ...
type PrimaryKey struct {
	key  string
	hash string
}

// ByPart provides comparator by partition type
type ByPart []*PrimaryKey

// Len ...
func (p ByPart) Len() int {
	return len(p)
}

// Less ...
func (p ByPart) Less(i, j int) bool {
	pType := config.HashingStrategy
	switch pType {
	case config.Random:
		return p[i].hash < p[j].hash
	case config.Ophf:
		return p[i].key < p[j].key
	default:
		return p[i].hash < p[j].hash
	}
}

// Swap ...
func (p ByPart) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func createPrimaryKeys(keys []string) []*PrimaryKey {
	res := make([]*PrimaryKey, 0)
	for _, key := range keys {
		res = append(res, NewPrimaryKey(key))
	}
	sort.Sort(ByPart(res))
	return res
}

// NewPrimaryKey ...
func NewPrimaryKey(key string) *PrimaryKey {
	p := &PrimaryKey{}
	p.key = key
	pType := config.HashingStrategy
	switch pType {
	case config.Random:
		p.hash = hash(p.key)
	default:
		p.hash = hash(p.key)
	}
	return p
}

func hash(key string) string {
	tmp := md5.Sum([]byte(key))
	return string(tmp[:])
}
