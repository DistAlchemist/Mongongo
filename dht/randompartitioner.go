// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package dht

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/DistAlchemist/Mongongo/config"
)

// RandomPartInstance ...
var (
	RandomPartInstance = NewRandomPartitioner()
	s                  = rand.NewSource(time.Now().UnixNano())
	rnd                = rand.New(s)
	sid, err           = os.Hostname()
)

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

// GetDefaultToken ...
func (r *RandomPartitioner) GetDefaultToken() string {
	initialToken := config.InitialToken
	if initialToken != "" {
		return initialToken
	}
	// generate random token
	guid := getGUID()
	token := r.Hash(guid)
	return token
}

func getGUID() string {
	t := time.Now().UnixNano() / int64(time.Millisecond)
	r := rnd.Int63()
	res := string(sid) + ":" +
		fmt.Sprint(t) + ":" +
		fmt.Sprint(r)
	return res
}

// Hash ...
func (r *RandomPartitioner) Hash(key string) string {
	tmp := md5.Sum([]byte(key))
	return string(tmp[:])
}
