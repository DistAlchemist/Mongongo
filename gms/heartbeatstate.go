// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

import "sync/atomic"

// HeartBeatState associated with any given endpoint.
type HeartBeatState struct {
	generation int
	heartBeat  int32 // atomic
	version    int32
}

// NewHeartBeatState creates a new hbState with given
// generation and hearBeat. verison defaults to 0.
func NewHeartBeatState(generation, heartBeat int) *HeartBeatState {
	h := &HeartBeatState{}
	h.generation = generation
	h.heartBeat = int32(heartBeat)
	h.version = 0
	return h
}

// UpdateHeartBeat increments generation and atomically
// increments version
func (h *HeartBeatState) UpdateHeartBeat() {
	atomic.AddInt32(&h.heartBeat, 1)
	h.version = GetNextVersion()
}

// GetHeartBeat returns heartbeat
func (h *HeartBeatState) GetHeartBeat() int32 {
	return h.heartBeat
}

// GetVersion returns version
func (h *HeartBeatState) GetVersion() int32 {
	return h.version
}
