// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

import "github.com/DistAlchemist/Mongongo/network"

// IEndPointStateChangeSubscriber provides an interface
// for endpoint state change subscribers
type IEndPointStateChangeSubscriber interface {
	OnChange(endpoint network.EndPoint, epState *EndPointState)
}
