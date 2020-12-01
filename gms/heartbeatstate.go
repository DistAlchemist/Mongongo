// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

// HeartBeatState associated with any given endpoint.
type HeartBeatState struct {
	generation int
	heartBeat  int // atomic
	version    int
}
