package mapper

import (
	"Envoy-Pilot/cmd/server/constant"
	"fmt"

	"github.com/gogo/protobuf/types"
)

type MapperStruct interface {
	GetResources(configJson string) ([]types.Any, error)
}

// GetMapperFor given topic
func GetMapperFor(topic string) MapperStruct {
	switch topic {
	case constant.SUBSCRIBE_CDS:
		return &ClusterMapper{}
	case constant.SUBSCRIBE_LDS:
		return &ListenerMapper{}
	case constant.SUBSCRIBE_RDS:
		return &RouteMapper{}
	case constant.SUBSCRIBE_EDS:
		return &EndpointMapper{}
	default:
		panic(fmt.Sprintf("No mapper found for type %s\n", topic))
	}
}
