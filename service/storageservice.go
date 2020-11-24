package service

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"

	"github.com/DistAlchemist/Mongongo/config"
	"github.com/DistAlchemist/Mongongo/db"
	"github.com/DistAlchemist/Mongongo/locator"
	"github.com/DistAlchemist/Mongongo/network"
)

// StorageService apply functions to storage layer
type StorageService struct {
	//
	uptime int64
	// storageLoadBalancer *StorageLoadBalancer
	endpointSnitch locator.EndPointSnitch
	tokenMetadata  locator.TokenMetadata
	nodePicker     locator.RackStrategy
}

var (
	mu       sync.Mutex
	instance *StorageService
)

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
	// ss.storageLoadBalancer = StorageLoadBalancer{ss}
	ss.endpointSnitch = locator.EndPointSnitch{}
	ss.tokenMetadata = locator.TokenMetadata{}
	if config.RackAware == true {
		ss.nodePicker = locator.RackAwareStrategy{}
	} else {
		ss.nodePicker = locator.RackUnawareStrategy{}
	}
}

func (ss *StorageService) getNStorageEndPointMap(key string) map[network.EndPoint]network.EndPoint {
	// dummy implementation
	// should hash the key to get EndPoint map.
	res := make(map[network.EndPoint]network.EndPoint)
	res[network.EndPoint{"localhost", "11111"}] = network.EndPoint{"localhost", "11111"}
	return res
}

func (ss *StorageService) Start() {
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
	l, e := net.Listen("tcp", "localhost:11111")
	fmt.Println("StorageService listening to localhost:11111")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	go http.Serve(l, mux)
}

func (ss *StorageService) DoRowMutation(args *db.RowMutationArgs, reply *db.RowMutationReply) error {
	//
	fmt.Println("enter DoRowMutation")
	reply.Result = "DoRowMutation success"
	return nil
}
