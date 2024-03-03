package wfi

import (
	"testing"
)

func TestTCPPort(t *testing.T) {
	if !TryPort("tcp", "google.com", "443") {
		t.Fatal()
	}
}

func TestLocalhostTCP(t *testing.T) {
	if !TryPort("tcp", "localhost", "3306") {
		t.Fatal()
	}
}

func TestUDPPort(t *testing.T) {
	if !TryPort("udp", "1.1.1.1", "53") {
		t.Fatal()
	}
}
