// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package network

import (
	"log"
	"os"
)

// EndPoint stores hostname, ip addr, port number etc.
type EndPoint struct {
	//
	HostName string
	Port     string
}

// NewEndPoint will return an EndPoint instance
// constructed by localhost name and given port
func NewEndPoint(port string) *EndPoint {
	e := &EndPoint{}
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	e.HostName = hostname
	e.Port = port
	return e
}

// NewEndPointH init an EndPoint given hostname and port
func NewEndPointH(hostname string, port string) *EndPoint {
	e := &EndPoint{}
	e.HostName = hostname
	e.Port = port
	return e
}
