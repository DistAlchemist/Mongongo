// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package locator

import "github.com/DistAlchemist/Mongongo/network"

// RackStrategy is the interface for rack strategy
// RackAwareStrategy and RackUnawareStrategy are two implementations
type RackStrategy interface {
	GetStorageEndPoints(token uint64) map[network.EndPoint]network.EndPoint
}
