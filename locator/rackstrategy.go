// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package locator

import (
	"log"
	"sort"

	"github.com/DistAlchemist/Mongongo/config"
	"github.com/DistAlchemist/Mongongo/gms"
	"github.com/DistAlchemist/Mongongo/network"
)

// IStrategy is the interface for rack strategy
type IStrategy interface {
	GetStorageEndPoints(token string) []network.EndPoint
	GetTokenEndPointMap() map[string]network.EndPoint
	GetToken(endPoint network.EndPoint) string
	GetReadStorageEndPoints(token string) map[network.EndPoint]bool
	// GetHintedStorageEndPoints(token *big.Int) map[network.EndPoint]network.EndPoint
}

// RackStrategy implements IStrategy, adds its own methods
type RackStrategy struct {
	I IStrategy
}

// GetHintedStorageEndPoints returns a hinted map.
// The key is the endpoint on which the data is being placed.
// The value is the endpoint which is in the top N.
// Currently it is the map of top N to live nodes.
func (r *RackStrategy) GetHintedStorageEndPoints(token string) map[network.EndPoint]network.EndPoint {
	topN := r.I.GetStorageEndPoints(token) // N is # of replicas, see config.ReplicationFactor
	m := make(map[network.EndPoint]network.EndPoint)
	liveList := make([]network.EndPoint, 0)
	for _, endPoint := range topN {
		if gms.GetFailureDetector().IsAlive(endPoint) {
			m[endPoint] = endPoint
			liveList = append(liveList, endPoint)
		} else {
			nxt, ok := r.getNextAvailableEndPoint(endPoint, topN, liveList)
			if !ok {
				m[nxt] = endPoint // map: alive -> origin
				liveList = append(liveList, nxt)
			} else {
				log.Printf("Unable to find a live endpoint, we might run out of live endpoints! dangerous!\n")
			}
		}
	}
	return m
}

func (r *RackStrategy) getNextAvailableEndPoint(startPoint network.EndPoint,
	topN []network.EndPoint, liveNodes []network.EndPoint) (network.EndPoint, bool) {
	tokenToEndPointMap := r.I.GetTokenEndPointMap()
	tokens := make([]string, 0, len(tokenToEndPointMap))
	for k := range tokenToEndPointMap {
		tokens = append(tokens, k)
	}
	sort.Strings(tokens)
	token := r.I.GetToken(startPoint)
	idx := sort.SearchStrings(tokens, token)
	totalNodes := len(tokens)
	if idx == totalNodes {
		idx = 0
	}
	startIdx := (idx + 1) % totalNodes
	var endPoint network.EndPoint
	flag := false
	for i, count := startIdx, 1; count < totalNodes; count, i = count+1, (i+1)%totalNodes {
		tmp := tokenToEndPointMap[tokens[i]]
		if gms.GetFailureDetector().IsAlive(tmp) && !contains(topN, tmp) &&
			!contains(liveNodes, tmp) {
			endPoint = tmp
			flag = true
			break
		}
	}
	return endPoint, flag
}

// This function change endpoints' port to storage port
func retrofitPorts(list []network.EndPoint) {
	for _, tmp := range list {
		tmp.Port = config.StoragePort
	}
}

// // GetStorageEndPoints return endpoints that correspond to a given token.
// func (r *RackStrategy) GetStorageEndPoints(token *big.Int) map[network.EndPoint]network.EndPoint {
// 	return make(map[network.EndPoint]network.EndPoint)
// }

// GetReadStorageEndPoints ...
func (r *RackStrategy) GetReadStorageEndPoints(token string) map[network.EndPoint]bool {
	// return map[network.EndPoint]bool{}
	return r.I.GetReadStorageEndPoints(token)
}
