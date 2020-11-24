package service

// StorageLoadBalancer keeps load information across the system.
// It registers itself with the Gossiper for load information which is
// the number of requests processed w.r.t distinct keys at an endpoint.
// Monitor load information at the interval of 5 minutes then do
// load balancing operations if necessary.
type StorageLoadBalancer struct {
	// storageService *StorageService
}
