// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/DistAlchemist/Mongongo/service"
)

func main() {
	mg := new(service.Mongongo)
	mg.Hostname = "localhost"
	mg.Port = 9160
	mg.Start()
	serv := rpc.NewServer()
	serv.Register(mg)
	// ===== workaround ==========
	oldMux := http.DefaultServeMux
	mux := http.NewServeMux()
	http.DefaultServeMux = mux
	// ===========================
	serv.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	// ===== workaround ==========
	http.DefaultServeMux = oldMux
	// ===========================
	l, e := net.Listen("tcp", "localhost:9160")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	go http.Serve(l, mux)
	for {
		// wait
	}
}
