package locator

import "github.com/DistAlchemist/Mongongo/network"

type RackStrategy interface {
	GetStorageEndPoints([]string) map[string][]network.EndPoint
}
