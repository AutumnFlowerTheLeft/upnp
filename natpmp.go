package upnp

import (
	"net"
)

type AddPortMapping struct {
	Getgateway         net.IP
	Protocol           string
	Port               int
	MappedExternalPort int
	TimeOut            int
}

type AddPortMappingResult struct {
	SecondsSinceStartOfEpoc      uint32
	InternalPort                 uint16
	MappedExternalPort           uint16
	PortMappingLifetimeInSeconds uint32
}

func NewOpenUPnP(u *AddPortMapping) *AddPortMapping {
	return u
}

func OpenPortMapping(u *AddPortMapping) {
	var Protocolcode byte
	switch u.Protocol {
	case "udp":
		Protocolcode = 1
	case "tcp":
		Protocolcode = 2
	default:
		Protocolcode = 0
	}
	msg := make([]byte, 12)
	// Version 0
	msg[0] = 0
	msg[1] = Protocolcode

	writeNetworkOrderUint16(msg[4:6], uint16(u.Port))
	writeNetworkOrderUint16(msg[6:8], uint16(u.MappedExternalPort))
	writeNetworkOrderUint32(msg[8:12], uint32(u.TimeOut))
	// func Apply()* to network.go
	// response, e := Apply(u, msg, 16)
	Apply(u, msg, 16)
	// resp = &AddPortMappingResult{}
	// resp.SecondsSinceStartOfEpoc = readNetworkOrderUint32(response[4:8])
	// resp.InternalPort = readNetworkOrderUint16(response[8:10])
	// resp.MappedExternalPort = readNetworkOrderUint16(response[10:12])
	// resp.PortMappingLifetimeInSeconds = readNetworkOrderUint32(response[12:16])
}
func writeNetworkOrderUint16(buf []byte, d uint16) {
	buf[0] = byte(d >> 8)
	buf[1] = byte(d)
}

func writeNetworkOrderUint32(buf []byte, d uint32) {
	buf[0] = byte(d >> 24)
	buf[1] = byte(d >> 16)
	buf[2] = byte(d >> 8)
	buf[3] = byte(d)
}

func readNetworkOrderUint16(buf []byte) uint16 {
	return (uint16(buf[0]) << 8) | uint16(buf[1])
}

func readNetworkOrderUint32(buf []byte) uint32 {
	return (uint32(buf[0]) << 24) | (uint32(buf[1]) << 16) | (uint32(buf[2]) << 8) | uint32(buf[3])
}

func (upnp *AddPortMapping) OpenTCPPort(Port int) {
	gwy, _ := Getgateway()

	response := &AddPortMapping{
		Getgateway:         gwy,
		Protocol:           "tcp",
		Port:               Port,
		MappedExternalPort: 0,
		TimeOut:            3600,
	}
	OpenPortMapping(response)
}

func (upnp *AddPortMapping) OpenUDPPort(Port int) {
	gwy, _ := Getgateway()

	response := &AddPortMapping{
		Getgateway:         gwy,
		Protocol:           "udp",
		Port:               Port,
		MappedExternalPort: 0,
		TimeOut:            3600,
	}
	OpenPortMapping(response)
}
