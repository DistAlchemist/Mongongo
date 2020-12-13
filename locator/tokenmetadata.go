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
	bootstrapNodes     map[string]network.EndPoint
}

// NewTokenMetadata ...
func NewTokenMetadata() *TokenMetadata {
	t := &TokenMetadata{}
	t.tokenToEndPointMap = make(map[string]network.EndPoint)
	t.endPointToTokenMap = make(map[network.EndPoint]string)
	t.bootstrapNodes = make(map[string]network.EndPoint)
	return t
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

// CloneBootstrapNodes ...
func (t *TokenMetadata) CloneBootstrapNodes() map[string]network.EndPoint {
	t.rwm.RLock()
	defer t.rwm.RUnlock()
	res := make(map[string]network.EndPoint)
	for k, v := range t.bootstrapNodes {
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

// IsKnownEndPoint ...
func (t *TokenMetadata) IsKnownEndPoint(ep *network.EndPoint) bool {
	t.rwm.RLock()
	defer t.rwm.RUnlock()
	return t.endPointToTokenMap[*ep] != ""
}

// Update ...
func (t *TokenMetadata) Update(token string, endpoint *network.EndPoint, bootstrapState bool) {
	t.rwm.Lock()
	defer t.rwm.Unlock()
	if bootstrapState {
		t.bootstrapNodes[token] = *endpoint
		t.Remove(endpoint)
	} else {
		delete(t.bootstrapNodes, token)
		oldToken := t.endPointToTokenMap[*endpoint]
		if oldToken != "" {
			delete(t.tokenToEndPointMap, oldToken)
		}
		t.tokenToEndPointMap[token] = *endpoint
		t.endPointToTokenMap[*endpoint] = token
	}
}

// Remove ...
func (t *TokenMetadata) Remove(endpoint *network.EndPoint) {
	t.rwm.Lock()
	defer t.rwm.Unlock()
	oldToken := t.endPointToTokenMap[*endpoint]
	if oldToken != "" {
		delete(t.tokenToEndPointMap, oldToken)
	}
	delete(t.endPointToTokenMap, *endpoint)
}
