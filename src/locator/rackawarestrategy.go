package locator

import "github.com/DistAlchemist/Mongongo/network"

type RackAwareStrategy struct {
	//
}

func (ras RackAwareStrategy) GetStorageEndPoints([]string) map[string][]network.EndPoint {
	return make(map[string][]network.EndPoint)
}
