package upnp

import (
	"fmt"
	"net"
	"time"
)

const NAT_PMP_PORT = 5351
const TRIES = 9
const INITIAL_MS = 300

func Apply(u *AddPortMapping, msg []byte, timeout time.Duration) (result []byte, err error) {
	gatewayUPnP := &net.UDPAddr{
		IP:   u.Getgateway,
		Port: NAT_PMP_PORT,
	}
	c, e := net.DialUDP("udp", nil, gatewayUPnP)
	if e != nil {
		return
	}
	defer c.Close()
	result = make([]byte, 16)
	var finalTimeout time.Time
	if timeout != 0 {
		finalTimeout = time.Now().Add(timeout)
	}
	needNewDeadline := true

	var tries uint
	// Init Tries count = 0, Max tries for 9
	// if Timeout, try again.
	for tries = 0; (tries < TRIES && finalTimeout.IsZero()) || time.Now().Before(finalTimeout); {
		// TimeOut try again
		if needNewDeadline {
			nextDeadline := time.Now().Add((INITIAL_MS << tries) * time.Millisecond)
			err = c.SetDeadline(minTime(nextDeadline, finalTimeout))
			if err != nil {
				return
			}
			needNewDeadline = false
		}
		_, err = c.Write(msg)
		if err != nil {
			return
		}
		var bytesRead int
		var remoteAddr *net.UDPAddr
		bytesRead, remoteAddr, err = c.ReadFromUDP(result)
		if err != nil {
			if err.(net.Error).Timeout() {
				tries++
				needNewDeadline = true
				continue
			}
			return
		}
		if !remoteAddr.IP.Equal(u.Getgateway) {
			continue
		}

		if bytesRead < len(result) {
			result = result[:bytesRead]
		}
		return
	}
	err = fmt.Errorf("Timed out trying to contact gateway")
	return
}

func minTime(a, b time.Time) time.Time {
	if a.IsZero() {
		return b
	}
	if b.IsZero() {
		return a
	}
	if a.Before(b) {
		return a
	}
	return b
}

func protocolChecks(msg []byte, resultSize int, result []byte) (err error) {
	if len(result) != resultSize {
		err = fmt.Errorf("unexpected result size %d, expected %d", len(result), resultSize)
		return
	}
	if result[0] != 0 {
		err = fmt.Errorf("Not support Protocol version %d", result[0])
		return
	}
	expectedOp := msg[1] | 0x80
	if result[1] != expectedOp {
		err = fmt.Errorf("Unexpected Protocol %d. Expected %d", result[1], expectedOp)
		return
	}
	resultCode := readNetworkOrderUint16(result[2:4])
	if resultCode != 0 {
		err = fmt.Errorf("Error result code %d", resultCode)
		return
	}
	return
}
