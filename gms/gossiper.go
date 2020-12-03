// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

import (
	"sort"
	"sync"
	"time"

	"github.com/DistAlchemist/Mongongo/config"
	"github.com/DistAlchemist/Mongongo/network"
)

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

	localEndPoint        *network.EndPoint
	aVeryLongTime        int64
	preIdx               int // index used previously
	rrIdx                int // round robin index through live endpoint set
	liveEndpoints        map[network.EndPoint]bool
	unreachableEndpoints map[network.EndPoint]bool
	seeds                map[network.EndPoint]bool
	endPointStateMap     map[network.EndPoint]*EndPointState
	subscribers          []IEndPointStateChangeSubscriber
	mu                   sync.Mutex
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
	g.intervalInMillis = 1000
	g.aVeryLongTime = 259200 * 1000
	g.preIdx = 0
	g.rrIdx = 0
	g.liveEndpoints = make(map[network.EndPoint]bool)
	g.unreachableEndpoints = make(map[network.EndPoint]bool)
	g.seeds = make(map[network.EndPoint]bool)
	g.endPointStateMap = make(map[network.EndPoint]*EndPointState)
	g.subscribers = make([]IEndPointStateChangeSubscriber, 0)
	GetFailureDetector().RegisterEventListener(g)
	return g
}

// GetGossiper creates a new Gossiper if not exists
func GetGossiper() *Gossiper {
	if gossiper == nil {
		gossiper = NewGossiper()
	}
	return gossiper
}

// Start will start gossiper on control port
func (g *Gossiper) Start(generation int) {
	g.localEndPoint = network.NewEndPoint(config.ControlPort)
	// get the seeds from the config and initialize them.
	seedHosts := config.Seeds
	for seedHost := range seedHosts {
		seed := network.NewEndPointH(seedHost, config.ControlPort)
		if *seed == *g.localEndPoint {
			// already this host
			continue
		}
		g.seeds[*seed] = true // add seed host
	}

	// initialize the heartbeat state for this localEndPoint
	localState, ok := g.endPointStateMap[*g.localEndPoint]
	if !ok {
		// localState doesn't exist
		hbState := NewHeartBeatState(generation, 0)
		localState = NewEndPointState(hbState)
		localState.SetAlive(true)
		localState.SetGossiper(true)
		g.endPointStateMap[*g.localEndPoint] = localState
	}
	go g.RunTimerTask()
}

// RunTimerTask starts the periodic task for a gossiper
func (g *Gossiper) RunTimerTask() {
	// currently it runs every 1 min
	for {
		g.runTask()
		time.Sleep(time.Millisecond * time.Duration(g.intervalInMillis))
	}
}

func (g *Gossiper) runTask() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.endPointStateMap[*g.localEndPoint].GetHeartBeatState().UpdateHeartBeat()
	gDigests := make([]*GossipDigest, 0)
	g.makeRandomGossipDigest(gDigests)
	if len(gDigests) > 0 {
		// TODO will send gossip messages to other nodes
	}
}

func (g *Gossiper) makeRandomGossipDigest(gDigests []*GossipDigest) {
	// the gossip digest is built based on randomization rather than
	// just looping through the collection of live endpoints.
	epState := g.endPointStateMap[*g.localEndPoint]
	generation := int(epState.GetHeartBeatState().generation)
	maxVersion := getMaxEndPointStateVersion(epState)
	gDigests = append(gDigests, NewGossipDigest(*g.localEndPoint, generation, maxVersion))
	// map is unsorted, so we omit the shuffle here
	for liveEndPoint := range g.liveEndpoints {
		epState, ok := g.endPointStateMap[liveEndPoint]
		if ok {
			generation = int(epState.GetHeartBeatState().generation)
			maxVersion = getMaxEndPointStateVersion(epState)
			gDigests = append(gDigests, NewGossipDigest(liveEndPoint, generation, maxVersion))
		} else {
			gDigests = append(gDigests, NewGossipDigest(liveEndPoint, 0, 0))
		}
	}
}

func getMaxEndPointStateVersion(epState *EndPointState) int {
	versions := make([]int, 0)
	versions = append(versions, int(epState.GetHeartBeatState().GetVersion()))
	appStateMap := epState.applicationState
	for key := range appStateMap {
		stateVersion := appStateMap[key].version
		versions = append(versions, int(stateVersion))
	}
	sort.Ints(versions)
	return versions[len(versions)-1]
}

// GetEndPointStateForEndPoint returns state for given endpoint.
func (g *Gossiper) GetEndPointStateForEndPoint(ep network.EndPoint) EndPointState {
	return *g.endPointStateMap[ep]
}

// Register register end point state change subscriber
func (g *Gossiper) Register(subscriber IEndPointStateChangeSubscriber) {
	g.subscribers = append(g.subscribers, subscriber)
}

// Convict implements IFailureDetectionEventListener interface
// it is invoked by the Failure Detector when it convicts an end point
func (g *Gossiper) Convict(endpoint network.EndPoint) {
	// TODO
}

// Suspect implements IFailureDetectionEventListener interface
// it is invoked by the Failure Detector when it suspects an end point
func (g *Gossiper) Suspect(endpoint network.EndPoint) {
	// TODO
}
