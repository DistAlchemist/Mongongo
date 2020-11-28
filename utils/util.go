package utils

// I'm not intended to provide a general function interaface currently for convenience.

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

// GetBytes Convert any type data into bytes
func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(key); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Hash Convert key into an integer
func Hash(key interface{}) uint64 {
	rawbytes, _ := GetBytes(key)
	h := fnv.New64()
	h.Write(rawbytes)
	return h.Sum64()
}
