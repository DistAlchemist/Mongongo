// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import "os"

// FileStruct ...
type FileStruct struct {
	key    string
	reader *os.File
	buf    []byte
}
