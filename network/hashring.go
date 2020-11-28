package network

// HashRing is response for finding coordinator of a given key

import (
	"errors"

	"github.com/DistAlchemist/Mongongo/config"
	"github.com/DistAlchemist/Mongongo/utils"
)

// HashRing manages the behavior of all actions on HashRing
// type HashRing interface {
// HashToEndPoint(hash int) EndPoint
//FindCoordinator(hash uint64) (*EndPoint, error)
//AddEndPoint(host EndPoint) error
//MoveEndPoint(host EndPoint, hash uint32) error
//RemoveEndPoint() error
//NextEndPoint() EndPoint
//}

// HostNode contains all information about single node in a HashRing
type HostNode struct {
	endPoint EndPoint
	location uint64
	preNode  *HostNode
	nextNode *HostNode
}

// HashRing is the consistent hash ring.
type HashRing struct {
	ringHead *HostNode // maybe we can use the built-in ring package
	maxRange uint64
	nodeNum  int
}

// CreateHashRing creates a new HashRing.
func CreateHashRing() *HashRing {
	return &HashRing{maxRange: uint64(config.RingRange), nodeNum: 0}
}

// FindCoordinator finds a host responsible for a given value.
func (ring *HashRing) FindCoordinator(data interface{}) (*EndPoint, error) {
	hashValue := (utils.Hash(data)) % ring.maxRange
	hostNode, err := ring.findHostNode(hashValue)
	if err != nil {
		return nil, err
	}
	return &hostNode.endPoint, nil
}

// AddNode adds a new HostNode into the HashRing.
func (ring *HashRing) AddNode(endPoint EndPoint) error {
	if uint64(ring.nodeNum) >= ring.maxRange {
		return errors.New("HASHRING: NO EXTRA SPACE FOR A NEW HOST TO INSERT")
	}

	hostLocation := utils.Hash(endPoint) % ring.maxRange

	// ensure the location is not occupied
	if ring.nodeNum == 0 {
		// no hostnode in the ring, so we can add a new host directly
		hostNode := HostNode{endPoint: endPoint, location: hostLocation}
		hostNode.preNode = &hostNode
		hostNode.nextNode = &hostNode

		ring.ringHead = &hostNode
	} else {
		possibleHost, _ := ring.findHostNode(hostLocation)
		// loop until finding a safe location
		for hostLocation == possibleHost.location {
			hostLocation = utils.Hash(hostLocation) % ring.maxRange
			possibleHost, _ = ring.findHostNode(hostLocation)
		}
		// insert the hostnode into the hashring
		hostNode := HostNode{endPoint: endPoint, location: hostLocation}
		hostNode.preNode = possibleHost.preNode
		hostNode.nextNode = possibleHost.nextNode
		hostNode.preNode.nextNode = &hostNode
		hostNode.nextNode.preNode = &hostNode

		if hostNode.location < ring.ringHead.location {
			ring.ringHead = &hostNode
		}
	}
	ring.nodeNum++
	return nil
}

// RemoveNode removes endpoint from the hashring
func (ring *HashRing) RemoveNode(endPoint EndPoint) {
	possibleHost := ring.endPointToHost(endPoint)
	if possibleHost != nil {
		possibleHost.preNode.nextNode = possibleHost.nextNode
		possibleHost.nextNode.preNode = possibleHost.preNode
	}
}

// PreNode returns the information of previous endpoint in the hashring
func (ring *HashRing) PreNode(endPoint EndPoint) *EndPoint {
	possibleHost := ring.endPointToHost(endPoint)
	if possibleHost == nil {
		return nil
	}
	return &possibleHost.preNode.endPoint
}

// NextNode returns the information of next endpoint in the hashring
func (ring *HashRing) NextNode(endPoint EndPoint) *EndPoint {
	possibleHost := ring.endPointToHost(endPoint)
	if possibleHost == nil {
		return nil
	}
	return &possibleHost.nextNode.endPoint
}

//-----------------------------------------------------------------
// private functions

func (ring *HashRing) endPointToHost(endPoint EndPoint) *HostNode {
	node := ring.ringHead
	for i := 0; i < ring.nodeNum; i++ {
		if node.endPoint == endPoint {
			return node
		}
		node = node.nextNode
	}
	return nil
}

func (ring *HashRing) findHostNode(hashValue uint64) (*HostNode, error) {
	if ring.nodeNum == 0 {
		return nil, errors.New("HASHRING: THERE IS NO HOSTNODE IN HASHRING")
	}
	if hashValue > ring.ringHead.preNode.location {
		// if hashValue is greater than the biggest hashvalue in the ring
		return ring.ringHead, nil
	} else {
		node := ring.ringHead
		for hashValue > node.location {
			node = node.nextNode
		}
		return node, nil
	}
}
