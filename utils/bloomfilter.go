// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utils

import (
	"encoding/binary"
	"log"
	"math/rand"
	"time"

	"github.com/willf/bitset"
)

// BloomFilter provides an approach to quickly
// check whether a key is stored inside some file
type BloomFilter struct {
	count  int
	size   int
	hashes int
	rnd    *rand.Rand
	filter *bitset.BitSet
}

var hashLibrary = []ISimpleHash{
	&RSHash{},
	&JSHash{},
	&PJWHash{},
	&ELFHash{},
	&BKDRHash{},
	&SDBMHash{},
	&DJBHash{},
	&DEKHash{},
	&BPHash{},
	&FNVHash{},
	&APHash{},
}

// NewBloomFilterS creates a bf deserialized from data
func NewBloomFilterS(count, hashes, bitsize int32, bs *bitset.BitSet) *BloomFilter {
	bf := &BloomFilter{}
	bf.count = int(count)
	bf.size = int(bitsize)
	bf.hashes = int(hashes)
	bf.filter = bs
	return bf
}

// NewBloomFilterDS creates a bf deserialized from data
func NewBloomFilterDS(hashes int32, filter *bitset.BitSet) *BloomFilter {
	bf := &BloomFilter{}
	bf.hashes = int(hashes)
	bf.filter = filter
	return bf
}

// NewBloomFilter gives out a new instance of bf
func NewBloomFilter(numElements, bitsPerElement int) *BloomFilter {
	if numElements < 0 || bitsPerElement < 1 {
		log.Fatal("# of elements and bits must be non-negative.")
	}
	b := &BloomFilter{}
	s := rand.NewSource(time.Now().UnixNano())
	b.rnd = rand.New(s)
	// add a small random number of bits so that even if the set
	// of elements hasn't changed, we'll get different false positives.
	b.count = numElements
	b.size = numElements*bitsPerElement + 20 + b.rnd.Intn(64)
	b.filter = bitset.New(uint(b.size))
	b.hashes = 8
	return b
}

// Fill insert a key to bf
func (b *BloomFilter) Fill(key string) {
	for i := 0; i < b.hashes; i++ {
		hashValue := hashLibrary[i].hash(key)
		idx := hashValue % b.size
		if idx < 0 {
			idx = (-1) * idx
		}
		b.filter.Set(uint(idx))
	}
}

// IsPresent checks whether a key exists
func (b *BloomFilter) IsPresent(key string) bool {
	res := true
	for i := 0; i < b.hashes; i++ {
		hashValue := hashLibrary[i].hash(key)
		idx := hashValue % b.size
		if idx < 0 {
			idx = (-1) * idx
		}
		if !b.filter.Test(uint(idx)) {
			res = false
			return res
		}
	}
	return res
}

// ToByteArray serializes this bf
func (b *BloomFilter) ToByteArray() []byte {
	buf := make([]byte, 0)
	// write out the count of the bloom filter
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(b.count))
	buf = append(buf, b4...)
	// write the number of hash function used
	binary.BigEndian.PutUint32(b4, uint32(b.hashes))
	buf = append(buf, b4...)
	// write the size of this bf
	binary.BigEndian.PutUint32(b4, uint32(b.size))
	buf = append(buf, b4...)
	// write BitSet bytes
	bs, err := b.filter.MarshalBinary()
	if err != nil {
		log.Print(err)
	}
	buf = append(buf, bs...)
	return buf
}

// ISimpleHash provides an interface for hash functions
type ISimpleHash interface {
	hash(str string) int
}

// RSHash ...
type RSHash struct{}

func (r *RSHash) hash(s string) int {
	b, a, h := 378551, 63689, 0
	for i := 0; i < len(s); i++ {
		h = h*a + int(s[i])
		a = a * b
	}
	return h
}

// JSHash ...
type JSHash struct{}

func (j *JSHash) hash(s string) int {
	h := 1315423911
	for i := 0; i < len(s); i++ {
		h ^= ((h << 5) + int(s[i]) + (h >> 2))
	}
	return h
}

// PJWHash ...
type PJWHash struct{}

func (p *PJWHash) hash(s string) int {
	bitsInUnsignedInt := 4 * 8
	threeQuarters := bitsInUnsignedInt * 3 / 4
	oneEighth := bitsInUnsignedInt / 8
	highBits := (0xFFFFFFFF) << (bitsInUnsignedInt - oneEighth)
	h, test := 0, 0
	for i := 0; i < len(s); i++ {
		h = (h << oneEighth) + int(s[i])
		test = h & highBits
		if test != 0 {
			h = ((h ^ (test >> threeQuarters)) & (^highBits))
		}
	}
	return h
}

// ELFHash ...
type ELFHash struct{}

func (e *ELFHash) hash(s string) int {
	h, x := 0, 0
	for i := 0; i < len(s); i++ {
		h = (h << 4) + int(s[i])
		x = h & 0xF0000000
		if x != 0 {
			h ^= (x >> 24)
		}
		h &= ^x
	}
	return h
}

// BKDRHash ...
type BKDRHash struct{}

func (b *BKDRHash) hash(s string) int {
	seed := 131 // 31 131 1313 13131 131313 etc.
	h := 0
	for i := 0; i < len(s); i++ {
		h = (h * seed) + int(s[i])
	}
	return h
}

// SDBMHash ...
type SDBMHash struct{}

func (d *SDBMHash) hash(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = int(s[i]) + (h << 6) + (h << 16) - h
	}
	return h
}

// DJBHash ...
type DJBHash struct{}

func (d *DJBHash) hash(s string) int {
	h := 5381
	for i := 0; i < len(s); i++ {
		h = (h << 5) + h + int(s[i])
	}
	return h
}

// DEKHash ...
type DEKHash struct{}

func (d *DEKHash) hash(s string) int {
	h := len(s)
	for i := 0; i < len(s); i++ {
		h = (h << 5) ^ (h >> 27) ^ int(s[i])
	}
	return h
}

// BPHash ...
type BPHash struct{}

func (b *BPHash) hash(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = (h << 7) ^ int(s[i])
	}
	return h
}

// FNVHash ...
type FNVHash struct{}

func (f *FNVHash) hash(s string) int {
	fnvPrime := 0x811C9DC5
	h := 0
	for i := 0; i < len(s); i++ {
		h *= fnvPrime
		h ^= int(s[i])
	}
	return h
}

// APHash ...
type APHash struct{}

func (a *APHash) hash(s string) int {
	h := 0xAAAAAAAA
	for i := 0; i < len(s); i++ {
		if (i & 1) == 0 {
			h ^= ((h << 7) ^ int(s[i]) ^ (h >> 3))
		} else {
			h ^= (^((h << 11) ^ int(s[i]) ^ (h >> 5)))
		}
	}
	return h
}
