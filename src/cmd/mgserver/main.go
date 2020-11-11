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
	mg.Port = 1111
	rpc.Register(mg)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "localhost:1111")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	go http.Serve(l, nil)
	for {
		// wait
	}
}
