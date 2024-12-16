package upnp

import (
	"testing"
)

func TestNatmap(t *testing.T) {
	gwy, e := Getgateway()
	if e != nil {
		t.Fatal(e)
	}
	response := &AddPortMapping{
		Getgateway:         gwy,
		Protocol:           "tcp",
		Port:               5005,
		MappedExternalPort: 0,
		TimeOut:            3600,
	}
	OpenPortMapping(response)
	// slog.Info("INFO", "SecondsSinceStartOfEpoc", resp.SecondsSinceStartOfEpoc, "InternalPort", resp.InternalPort, "MappedExternalPort", resp.MappedExternalPort, "PortMappingLifetimeInSeconds", resp.PortMappingLifetimeInSeconds)
}
