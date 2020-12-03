// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package gms

import (
	"log"
	"os"
	"time"

	"github.com/DistAlchemist/Mongongo/config"
	"github.com/DistAlchemist/Mongongo/network"
)

var failureDetector IFailureDetector

// FailureDetector implements IFailureDetector
type FailureDetector struct {
	sampleSize      int
	phiSuspectThres int
	phiConvictThres int
	// Failure Detector has to have been up for at least
	// 1 min.
	uptimeThres int64
	// Time when the module was instantiated.
	creationTime     int64
	fdEventListeners []IFailureDetectionEventListener
}

// GetFailureDetector will create a new instance
// of FailureDetector if not exists
func GetFailureDetector() IFailureDetector {
	if failureDetector == nil {
		failureDetector = newFailureDetector()
	}
	return failureDetector
}

func newFailureDetector() *FailureDetector {
	f := &FailureDetector{}
	f.creationTime = time.Now().UnixNano() / int64(time.Millisecond)
	f.sampleSize = 1000
	f.phiSuspectThres = 5
	f.phiConvictThres = 8
	f.uptimeThres = 60000 // 1 min.
	return f
}

// IsAlive check whether the endpoint is up.
func (f *FailureDetector) IsAlive(ep network.EndPoint) bool {
	localHost, err := os.Hostname()
	if err != nil {
		log.Fatalf("error when getting hostname: %v\n", err)
	}
	if localHost == ep.HostName {
		return true
	}
	ep2 := network.EndPoint{HostName: ep.HostName, Port: config.ControlPort}
	epState := GetGossiper().GetEndPointStateForEndPoint(ep2)
	return epState.IsAlive()
}

// RegisterEventListener registers event listener for fd
func (f *FailureDetector) RegisterEventListener(listener IFailureDetectionEventListener) {
	f.fdEventListeners = append(f.fdEventListeners, listener)
}
