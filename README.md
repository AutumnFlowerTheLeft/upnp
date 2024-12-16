# Go-NAT-PMP
A Go language for the NAT-PMP internet protocol for port mapping and discovering the external IP address.
but this is XoRPC plugin, Open Transport.

Qing Connect RPC by Xianglake Studio

XoRPC : https://github.com/qiaoliangXgamemode/XoRPC
# Get the package

```
go get github.com/AutumnFlowerTheLeft/upnp@v1.0.0
```

# Usage
Create a file main.go with contents:
```
package main

import (
	"github.com/AutumnFlowerTheLeft/upnp"
)

func main() {
	gwy, e := upnp.Getgateway()
	if e != nil {
		panic(e)
	}
	response := &upnp.AddPortMapping{
		Getgateway:         gwy,
		Protocol:           "tcp",
		Port:               5005,
		MappedExternalPort: 0,
		TimeOut:            3600,
	}
	upnp.OpenPortMapping(response)
}
```

# Example
## Getgateway
Route Getgateway
```
gwy, e := upnp.Getgateway()
```


## AddPortMapping
```
Getgateway 
Protocol => tcp/udp
Port
MappedExternalPort
TimeOut
```