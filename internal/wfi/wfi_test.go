package wfi

import (
	"testing"
)

func TestTryPort(t *testing.T) {
	testCases := []struct {
		name  string
		proto string
		host  string
		port  string
		want  bool
	}{
		{"ValidTCPPort", "tcp", "google.com", "443", true},
		{"LocalTCPPort", "tcp", "localhost", "3306", true},
		{"ValidUDPPort", "udp", "1.1.1.1", "53", true},
		{"InvalidHostname", "tcp", "souzej.pk", "239", false},
		{"InvalidTCPPort", "tcp", "facebook.com", "92821", false},
		{"InvalidUDPPort", "udp", "github.com", "99999", false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := TryPort(tc.proto, tc.host, tc.port)
			if got != tc.want {
				t.Errorf("TryPort(%s, %s, %s) = %v; want %v", tc.proto, tc.host, tc.port, got, tc.want)
			}
		})
	}
}
