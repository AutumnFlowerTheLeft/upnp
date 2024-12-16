package upnp

import (
	"net"

	"github.com/jackpal/gateway"
)

func Getgateway() (net.IP, error) {
	gatewayIP, e := gateway.DiscoverGateway()
	if e != nil {
		return nil, e
	}
	return gatewayIP, nil
}
