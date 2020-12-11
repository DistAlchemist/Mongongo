// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

// ApplicationState is the state associated with a
// particular node which an application wants to
// make available to the rest of the nodes in the
// cluster.
// Whenever a piece of state needs to be disseminated
// to the rest of cluster, wrap the state in an instance
// of ApplicationState and add it to the Gossiper.
// e.g. if we want to disseminate load information for
// node A, do the following:
//   loadState := ApplicationState{<load string>}
//   gms.GetGossiper().AddAppState("LOAD STATE", loadState)
type ApplicationState struct {
	version int
	state   string
}

// NewApplicationState ...
func NewApplicationState(state string, version int) *ApplicationState {
	p := &ApplicationState{}
	p.version = version
	p.state = state
	return p
}

// NewApplicationStateS ...
func NewApplicationStateS(state string) *ApplicationState {
	p := &ApplicationState{}
	p.state = state
	p.version = int(GetNextVersion())
	return p
}

// GetState ...
func (p *ApplicationState) GetState() string {
	return p.state
}

// GetStateVersion ...
func (p *ApplicationState) GetStateVersion() int {
	return p.version
}
