// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package locator

import (
	"sync"

	"github.com/DistAlchemist/Mongongo/network"
)

// TokenMetadata matains information of tokens
type TokenMetadata struct {
	//
	rwm                sync.RWMutex
	tokenToEndPointMap map[string]network.EndPoint
	endPointToTokenMap map[network.EndPoint]string
}

// CloneTokenEndPointMap return a copy of current tokenToEndPointMap
func (t *TokenMetadata) CloneTokenEndPointMap() map[string]network.EndPoint {
	t.rwm.RLock()
	defer t.rwm.RUnlock()
	res := make(map[string]network.EndPoint, len(t.tokenToEndPointMap))
	for k, v := range t.tokenToEndPointMap {
		res[k] = v
	}
	return res
}

// GetToken retruns the corresponding token given endpoint.
func (t *TokenMetadata) GetToken(endpoint network.EndPoint) string {
	t.rwm.RLock()
	defer t.rwm.RUnlock()
	return t.endPointToTokenMap[endpoint]
}
