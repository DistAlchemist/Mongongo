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
	"github.com/DistAlchemist/Mongongo/dht"
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
	partitioner         dht.IPartitioner
	storageMetadata     *db.StorageMetadata
	isBootstrapMode     bool
	tcpAddr             *network.EndPoint
	udpAddr             *network.EndPoint
}

var (
	mu              sync.Mutex
	instance        *StorageService
	ssNodeID        = "NODE-IDENTIFIER"
	ssBootstrapMode = "BOOTSTRAP-MODE"
	ssInitialDelay  = 60000
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
	bootstrap := os.Getenv("bootstrap")
	ss.isBootstrapMode = bootstrap == "true"
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
	token := ss.partitioner.GetToken(key)
	return ss.nodePicker.GetHintedStorageEndPoints(token)
}

func (ss *StorageService) initPartitioner() {
	hashingStrategy := config.HashingStrategy
	if hashingStrategy == config.Ophf {
		ss.partitioner = dht.NewOPP()
	} else {
		ss.partitioner = dht.NewRandomPartitioner()
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
	ss.tcpAddr = network.NewEndPoint(config.StoragePort)
	ss.udpAddr = network.NewEndPoint(config.ControlPort)
	// _ = db.GetManagerInstance().Start()
	ss.startStorageServer()
	// ss.startControlServer()
	ss.storageLoadBalancer.start()
	gms.GetGossiper().Register(ss)
	gms.GetGossiper().Start(ss.storageMetadata.GetGeneration())
	// make sure this token gets gossiped around
	ss.tokenMetadata.Update(ss.storageMetadata.StorageID, network.NewEndPoint(config.StoragePort), ss.isBootstrapMode)
	state := gms.NewApplicationStateS(ss.storageMetadata.StorageID)
	gms.GetGossiper().AddApplicationState(ssNodeID, state)
	if ss.isBootstrapMode {
		log.Printf("starting in bootstrap mode\n")
		go ss.runBootStrap([]*network.EndPoint{ss.tcpAddr}, ss.storageMetadata.StorageID)
		gms.GetGossiper().AddApplicationState(ssBootstrapMode, gms.NewApplicationStateS(""))
	}
}

func (ss *StorageService) runBootStrap(targets []*network.EndPoint, tokens ...string) {
	// initial delay waiting for this node to get a stable endpoint map
	// defaults to 60s
	// time.Sleep(time.Duration(ssInitialDelay) * time.Millisecond)
	// // clone again now so we include all discovered nodes in out calculations
	// tokenMetadata := ss.tokenMetadata
	// mark as not bootstrapping to calculate ranges correctly
	// for i := 0; i < len(targets); i++ {
	// 	ss.tokenMetadata.Update(tokens[i], targets[i], false)
	// }
	// rangesWithSourceTarget := ss.getRangeWithSourceTarget()
	// log.Printf("beginning bootstrap process for %v ...\n", targets)
	// send messages to respective folks to stream data over to the
	// new nodes being bootstrapped
	// TODO
}

// func (ss *StorageService) getRangeWithSourceTarget() map[dht.Range]

// DoRowMutation as a rpc served by storage service
func (ss *StorageService) DoRowMutation(args *db.RowMutationArgs, reply *db.RowMutationReply) error {
	fmt.Println("enter DoRowMutation")
	db.DoRowMutation(args, reply)
	// apply row mutation
	// args.RM.Apply(db.NewRow(args.RM.RowKey))
	// reply.Result = "DoRowMutation success"
	return nil
}

func (ss *StorageService) getReadStorageEndPoints(key string) map[network.EndPoint]bool {
	return ss.nodePicker.GetReadStorageEndPoints(ss.partitioner.GetToken(key))
}

func (ss *StorageService) getLiveReadStorageEndPoints(key string) []network.EndPoint {
	// this method attemps to return N endpoints that are responsible
	// for storing the specified key for replication
	liveEps := make([]network.EndPoint, 0)
	endpoints := ss.getReadStorageEndPoints(key)
	for endpoint := range endpoints {
		if gms.GetFailureDetector().IsAlive(endpoint) {
			liveEps = append(liveEps, endpoint)
		}
	}
	return liveEps
}

func (ss *StorageService) deliverHints(endpoint *network.EndPoint) {
	db.GetHintedHandOffManagerInstance()
}

// OnChange implements interface for endpoint
// state change subscriber
func (ss *StorageService) OnChange(endpoint network.EndPoint, epState *gms.EndPointState) {
	// Called when there is a change in application state.
	// In particular we are interested in new tokens as a
	// result of a new node or an existing node moving to
	// a new location on the ring.
	ep := network.NewEndPointH(endpoint.HostName, config.StoragePort)
	// node identifier for this endpoint on the identifier space
	nodeIDState := epState.GetApplicationState(ssNodeID)
	// check if this has a bootstrapping state message
	bootstrapState := epState.GetApplicationState(ssBootstrapMode) != nil
	if bootstrapState {
		log.Printf("%v is in bootstrap state\n", ep.HostName)
	}
	if nodeIDState != nil {
		newToken := nodeIDState.GetState()
		log.Printf("change in state for %v - has token %v\n", endpoint, newToken)
		oldToken := ss.tokenMetadata.GetToken(*ep)
		if oldToken != "" {
			// if oldToken equals the newToken then the node
			// had crashed and is coming back up again. If oldToken
			// is not equal to the newToken this means that
			// the node is being relocated to another position
			// in the ring.
			if oldToken != newToken {
				log.Printf("relocation for endpoint: %v\n", ep)
				ss.tokenMetadata.Update(newToken, ep, bootstrapState)
			} else {
				// this means the node crashed and is coming back
				// up. deliver the hints that we have for this
				// endpoint
				log.Printf("sending hinted data to %v\n", ep)
				ss.deliverHints(&endpoint)
			}
		} else {
			// this is a new node and we just update the token map
			ss.tokenMetadata.Update(newToken, ep, bootstrapState)
		}
	} else {
		// if we are here and if this node is up and already has
		// an entry in the token map. it means that the node was
		// behind a network partition
		if epState.IsAlive() && ss.tokenMetadata.IsKnownEndPoint(&endpoint) {
			log.Printf("endpoint %v just recovered from a partition. sending hinted data\n",
				ep)
			ss.deliverHints(ep)
		}
	}
}

func (ss *StorageService) getHintedStorageEndpointMap(key string) map[network.EndPoint]network.EndPoint {
	return ss.nodePicker.GetHintedStorageEndPoints(ss.partitioner.GetToken(key))
}

func (ss *StorageService) doConsistencyCheck(row *db.Row, endpoints []network.EndPoint, command db.ReadCommand) {
	// TODO
	// go runConsistency
}
