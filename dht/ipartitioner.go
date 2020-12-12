// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package dht

// IPartitioner ...
type IPartitioner interface {
	DecorateKey(key string) string
	UndecorateKey(decoreateddKey string) string
	Compare(s1, s2 string) int
	GetToken(key string) string
}
