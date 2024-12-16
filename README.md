# Go-NAT-PMP
A Go language for the NAT-PMP internet protocol for port mapping and discovering the external IP address.

# Get the package

```
go get -u github.com/AutumnFlowerTheLeft/upnp
```

# Usage
Create a file main.go with contents:
```
package main

import (
	"testing"
)

func TestNatmap(t *testing.T) {
	gwy, e := Getgateway()
	if e != nil {
		t.Fatal(e)
	}
	response := &AddPortMapping{
		gateway:            gwy,
		Protocol:           "tcp",
		Port:               5005,
		MappedExternalPort: 0,
		TimeOut:            3600,
	}
	OpenPortMapping(response)
	// slog.Info("INFO", "SecondsSinceStartOfEpoc", resp.SecondsSinceStartOfEpoc, "InternalPort", resp.InternalPort, "MappedExternalPort", resp.MappedExternalPort, "PortMappingLifetimeInSeconds", resp.PortMappingLifetimeInSeconds)
}

```
