package upnp

import (
	"log/slog"
	"testing"
)

func TestRoute(T *testing.T) {
	gatewayIP, e := Getgateway()
	if e != nil {
		T.Error(e)
	}
	slog.Info("My gateway IP", "gatewayIP", gatewayIP)
}
