// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utils

import (
	"log"
	"os"

	"github.com/willf/bitset"
)

var (
	// BFSerializer ...
	BFSerializer = NewBloomFilterSerializer()
)

// BloomFilterSerializer ...
type BloomFilterSerializer struct{}

// NewBloomFilterSerializer ...
func NewBloomFilterSerializer() *BloomFilterSerializer {
	b := &BloomFilterSerializer{}
	return b
}

// Serialize ...
func (b *BloomFilterSerializer) Serialize(bf *BloomFilter, dos *os.File) {
	writeInt32(dos, int32(bf.hashes))
	bs, err := bf.filter.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	dos.Write(bs)
}

// SerializeB serialize bloom filter to byte slice
func (b *BloomFilterSerializer) SerializeB(bf *BloomFilter, dos []byte) {
	writeInt32B(dos, int32(bf.hashes))
	bs, err := bf.filter.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	dos = append(dos, bs...)
}

// Deserialize ...
func (b *BloomFilterSerializer) Deserialize(dis *os.File) *BloomFilter {
	hashes := readInt32(dis)
	var bs *bitset.BitSet
	bs.UnmarshalBinary(restBytes(dis))
	return NewBloomFilterDS(hashes, bs)
}
