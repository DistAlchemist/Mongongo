package locator

import "github.com/DistAlchemist/Mongongo/network"

type RackUnawareStrategy struct {
	//
}

func (rus RackUnawareStrategy) GetStorageEndPoints([]string) map[string][]network.EndPoint {
	return make(map[string][]network.EndPoint)
}
