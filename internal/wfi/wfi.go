package wfi

import (
	"net"
	"time"
)

func TryPort(proto string, host string, port string) bool {
	timeout := time.Second
	cnx, err := net.DialTimeout(proto, net.JoinHostPort(host, port), timeout)
	if err != nil {
		return false
	}
	if cnx != nil {
		defer cnx.Close()
		return true
	}
	return false
}
