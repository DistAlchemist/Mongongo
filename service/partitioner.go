// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

// import (
// 	"crypto/md5"
// 	"math/big"
//
//
//
//
//
//
//
//
//
// )

// // IPartitioner specifies the interface of a partitoner type
// type IPartitioner interface {
// 	hash(key string) string
// }

// type orderPreservingHashPartitioner struct {
// 	maxKeyHashLength int
// 	prime            *big.Int
// }

// // NewOrderPreservingHashPartitioner initializes and return
// // an orderPreservingHashPartitioner
// func NewOrderPreservingHashPartitioner() IPartitioner {
// 	o := orderPreservingHashPartitioner{}
// 	o.maxKeyHashLength = 24
// 	o.prime = big.NewInt(31)
// 	return o
// }

// func (p orderPreservingHashPartitioner) hash(key string) string {
// 	res := big.NewInt(0)
// 	len := len(key)
// 	var tmp *big.Int
// 	for i := 0; i < p.maxKeyHashLength; i++ {
// 		if i < len { // res = res * prime + key[i]
// 			res.Mul(res, p.prime)
// 			tmp.SetString(string(key[i]), 10)
// 			res.Add(res, tmp)
// 		} else { // res = res * prime + prime
// 			res.Mul(res, p.prime)
// 			res.Add(res, p.prime)
// 		}
// 	}
// 	return res.String()
// }

// type randomPartitioner struct{}

// // NewRandomPartitioner return a random partitioner which uses MD5 hash algo
// func NewRandomPartitioner() IPartitioner {
// 	r := randomPartitioner{}
// 	return r
// }

// func (p randomPartitioner) hash(key string) string {
// 	tmp := md5.Sum([]byte(key))
// 	return string(tmp[:])
// }
