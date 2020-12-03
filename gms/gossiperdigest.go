// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

import "github.com/DistAlchemist/Mongongo/network"

// GossipDigest contains information about a
// specified list of EndPoints and the largest
// version of the state they have generated as
// known by the local endpoint.
type GossipDigest struct {
	endPoint   network.EndPoint
	generation int
	maxVersion int
}

// NewGossipDigest creates a new gossip digest
func NewGossipDigest(endPoint network.EndPoint, generation, maxVersion int) *GossipDigest {
	g := &GossipDigest{}
	g.endPoint = endPoint
	g.generation = generation
	g.maxVersion = maxVersion
	return g
}
