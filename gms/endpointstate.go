// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

// EndPointState contains the HeartBeatState and
// ApplicationState.
type EndPointState struct {
	hbState          HeartBeatState
	applicationState map[string]ApplicationState
	updateTimestamp  int64
	isAlive          bool
	isAGossiper      bool
}

// IsAlive return liveiness state of this endpoint.
func (e *EndPointState) IsAlive() bool {
	return e.isAlive
}
