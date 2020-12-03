// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

// StorageLoadBalancer keeps load information across the system.
// It registers itself with the Gossiper for load information which is
// the number of requests processed w.r.t distinct keys at an endpoint.
// Monitor load information at the interval of 5 minutes then do
// load balancing operations if necessary.
type StorageLoadBalancer struct {
	storageService *StorageService
}

// NewStorageLoadBalancer initializes a storage load balancer.
func NewStorageLoadBalancer(ss *StorageService) *StorageLoadBalancer {
	slb := new(StorageLoadBalancer)
	slb.storageService = ss
	// StageManager.registerStage
	// MessagingService.registerVerbHandlers
	// storageService.registerComponentForShutdown
	return slb
}

// Start starts storage load balancer
func (s *StorageLoadBalancer) start() {
	// TODO
}
