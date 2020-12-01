// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

import "github.com/DistAlchemist/Mongongo/network"

var gossiper *Gossiper

// Gossiper is responsible for Gossiping information for
// the local endpoint. It maintains the list of live and
// dead endpoints. It will periodically (every 1 sec.)
// chooses a random node and initiates a round of Gossip
// with it.
// A round of Gossip involves 3 rounds of messaging:
// If A wants to initiate a round of Gossip with B:
//  1. A -> B using GossipDigestSynMessage.
//  2. B -> A using GossipDigestAckMessage.
//  3. A -> B using GossipDigestAck2Message.
// When this module heads from one of the above messages,
// it will update the FailureDetector with the liveness
// information.
type Gossiper struct {
	MaxGossipPacketSize  int
	GossipStage          string
	JoinVerbHandler      string
	GossipDigestSynVerb  string
	GossipDigestAckVerb  string
	GossipDigestAck2Verb string
	intervalInMillis     int

	localEndPoint        network.EndPoint
	aVeryLongTime        int64
	preIdx               int // index used previously
	rrIdx                int // round robin index through live endpoint set
	liveEndpoints        map[network.EndPoint]bool
	unreachableEndpoints map[network.EndPoint]bool
	seeds                map[network.EndPoint]bool
	endPointStateMap     map[network.EndPoint]EndPointState
}

// NewGossiper creates a new Gossiper
func NewGossiper() *Gossiper {
	g := &Gossiper{}
	g.MaxGossipPacketSize = 1428
	g.GossipStage = "GS" // abbr for Gossip Stage
	g.JoinVerbHandler = "JVH"
	g.GossipDigestSynVerb = "GSV"
	g.GossipDigestAckVerb = "GAV"
	g.GossipDigestAck2Verb = "GA2V"
	return g
}

// GetGossiper creates a new Gossiper if not exists
func GetGossiper() *Gossiper {
	if gossiper == nil {
		gossiper = NewGossiper()
	}
	return gossiper
}

// GetEndPointStateForEndPoint returns state for given endpoint.
func (g *Gossiper) GetEndPointStateForEndPoint(ep network.EndPoint) EndPointState {
	return g.endPointStateMap[ep]
}
