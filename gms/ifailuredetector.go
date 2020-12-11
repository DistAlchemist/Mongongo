// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

import "github.com/DistAlchemist/Mongongo/network"

// IFailureDetector provides an interface that can
// query liveness information of a node in the cluster.
type IFailureDetector interface {
	IsAlive(ep network.EndPoint) bool
	interpret(ep network.EndPoint)
	report(ep network.EndPoint)
	RegisterEventListener(listener IFailureDetectionEventListener)
	UnregisterEventListener(listener IFailureDetectionEventListener)
}
