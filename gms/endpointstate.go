// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

import "time"

// EndPointState contains the HeartBeatState and
// ApplicationState.
type EndPointState struct {
	hbState          *HeartBeatState
	applicationState map[string]*ApplicationState
	updateTimestamp  int64
	isAlive          bool
	isAGossiper      bool
}

// IsAlive return liveiness state of this endpoint.
func (e *EndPointState) IsAlive() bool {
	return e.isAlive
}

// NewEndPointState creates a new endpoint state
func NewEndPointState(hbState *HeartBeatState) *EndPointState {
	e := &EndPointState{}
	e.hbState = hbState
	e.applicationState = make(map[string]*ApplicationState)
	e.updateTimestamp = time.Now().UnixNano() / int64(time.Millisecond)
	e.isAlive = true
	e.isAGossiper = false
	return e
}

// SetAlive sets liveiness of the endpoint state
func (e *EndPointState) SetAlive(live bool) {
	e.isAlive = live
}

// SetGossiper sets whether it is a gossiper
func (e *EndPointState) SetGossiper(g bool) {
	e.isAGossiper = g
}

// GetHeartBeatState return hbState
func (e *EndPointState) GetHeartBeatState() *HeartBeatState {
	return e.hbState
}

// GetApplicationState ...
func (e *EndPointState) GetApplicationState(key string) *ApplicationState {
	return e.applicationState[key]
}

func (e *EndPointState) AddApplicationState(key string, appState *ApplicationState) {
	e.applicationState[key] = appState
}
