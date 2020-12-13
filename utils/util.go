// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utils

// I'm not intended to provide a general function interaface currently for convenience.

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// Logger ...
var (
	Logger *log.Logger
	lmu    sync.Mutex
)

// LoggerInstance ...
func LoggerInstance() *log.Logger {
	lmu.Lock()
	defer lmu.Unlock()
	if Logger == nil {
		timestamp := strconv.Itoa(int(CurrentTimeMillis()))
		path := "logs"
		name := path + string(os.PathSeparator) + "serverlog_" + timestamp
		err := os.MkdirAll(path, 0700)
		if err != nil {
			log.Print(err)
		}
		file, err := os.Create(name)
		if err != nil {
			log.Print(err)
		}
		Logger = log.New(file, "logger> ", 0)
	}
	return Logger
}

// CurrentTimeMillis ...
func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

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
