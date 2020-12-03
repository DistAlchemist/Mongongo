// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

import (
	"github.com/DistAlchemist/Mongongo/network"
)

// IFailureDetectionEventListener provides an interface
// for fd event listener
type IFailureDetectionEventListener interface {
	// convict the specified endpoint
	Convict(ep network.EndPoint)
	// suspect the specified endpoint
	Suspect(ep network.EndPoint)
}
