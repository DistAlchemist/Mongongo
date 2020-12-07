// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
	"time"

	"github.com/DistAlchemist/Mongongo/config"
	"github.com/DistAlchemist/Mongongo/db"
	"github.com/DistAlchemist/Mongongo/gms"
	"github.com/DistAlchemist/Mongongo/locator"
	"github.com/DistAlchemist/Mongongo/network"
)

// StorageService apply functions to storage layer
type StorageService struct {
	uptime              int64
	storageLoadBalancer *StorageLoadBalancer
	endpointSnitch      locator.EndPointSnitch
	tokenMetadata       locator.TokenMetadata
	nodePicker          *locator.RackStrategy
	partitioner         IPartitioner
	storageMetadata     *db.StorageMetadata
}

var (
	mu       sync.Mutex
	instance *StorageService
)

// GetInstance return storageServer instance
func GetInstance() *StorageService {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &StorageService{}
		instance.init()
	}
	return instance
}

func (ss *StorageService) init() {
	ss.uptime = time.Now().UnixNano() / int64(time.Millisecond)
	ss.storageLoadBalancer = NewStorageLoadBalancer(ss)
	ss.endpointSnitch = locator.EndPointSnitch{}
	ss.tokenMetadata = locator.TokenMetadata{}
	if config.RackAware == true {
		ss.nodePicker = &locator.RackStrategy{I: &locator.RackAwareStrategy{}} // locator.RackAwareStrategy{}
	} else {
		ss.nodePicker = &locator.RackStrategy{I: &locator.RackUnawareStrategy{}} // locator.RackUnawareStrategy{}
	}
}

func (ss *StorageService) getNStorageEndPointMap(key string) map[network.EndPoint]network.EndPoint {
	token := ss.partitioner.hash(key)
	return ss.nodePicker.GetHintedStorageEndPoints(token)
}

func (ss *StorageService) initPartitioner() {
	hashingStrategy := config.HashingStrategy
	if hashingStrategy == config.Ophf {
		ss.partitioner = NewOrderPreservingHashPartitioner()
	} else {
		ss.partitioner = NewRandomPartitioner()
	}
}

func (ss *StorageService) startStorageServer() {
	serv := rpc.NewServer()
	serv.Register(ss)
	// ===== workaround ==========
	oldMux := http.DefaultServeMux
	mux := http.NewServeMux()
	http.DefaultServeMux = mux
	// ===========================
	serv.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	// ===== workaround ==========
	http.DefaultServeMux = oldMux
	// ===========================
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	addr := hostname + ":" + config.StoragePort
	l, e := net.Listen("tcp", addr)
	log.Printf("StorageServer listening to %v\n", addr)
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	go http.Serve(l, mux)
}

func (ss *StorageService) startControlServer() {
	serv := rpc.NewServer()
	serv.Register(ss)
	// ===== workaround ==========
	oldMux := http.DefaultServeMux
	mux := http.NewServeMux()
	http.DefaultServeMux = mux
	// ===========================
	serv.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	// ===== workaround ==========
	http.DefaultServeMux = oldMux
	// ===========================
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	addr := hostname + ":" + config.ControlPort
	l, e := net.Listen("udp", addr)
	log.Printf("ControlServer listening to %v\n", addr)
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	go http.Serve(l, mux)
}

// Start will setup RPC server for storage service
func (ss *StorageService) Start() {
	ss.initPartitioner()
	ss.storageMetadata = db.GetManagerInstance().Start()
	// _ = db.GetManagerInstance().Start()
	ss.startStorageServer()
	// ss.startControlServer()
	ss.storageLoadBalancer.start()
	gms.GetGossiper().Register(ss)
	gms.GetGossiper().Start(ss.storageMetadata.GetGeneration())
	// make sure this token gets gossiped around
	// TODO
	// ss.tokenMetadata.Update(ss.storageMetadata.storageID)
	// gms.GetGossiper().AddApplicationState(..)
}

// DoRowMutation as a rpc served by storage service
func (ss *StorageService) DoRowMutation(args *db.RowMutationArgs, reply *db.RowMutationReply) error {
	fmt.Println("enter DoRowMutation")
	// TODO check hints
	// apply row mutation
	args.RM.Apply(db.NewRow(args.RM.RowKey))
	reply.Result = "DoRowMutation success"
	return nil
}

// OnChange implements interface for endpoint
// state change subscriber
func (ss *StorageService) OnChange(endpoint network.EndPoint, epState gms.EndPointState) {
	// TODO
}
