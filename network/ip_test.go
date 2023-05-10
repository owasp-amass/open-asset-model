package network_test

import (
	"encoding/json"
	"fmt"
	"net/netip"
	"testing"

	. "github.com/owasp-amass/open-asset-model/network"

	"github.com/stretchr/testify/require"
)

func TestIPAddress(t *testing.T) {
	tests := []struct {
		description string
		ip          string
		netType     string
	}{
		{
			description: "Test successful creation of IPAddress with valid IPv4 address",
			ip:          "192.168.1.1",
			netType:     "IPv4",
		},
		{
			description: "Test successful creation of IPAddress with valid IPv6 address",
			ip:          "2001:db8::1",
			netType:     "IPv6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			ip, _ := netip.ParseAddr(tt.ip)
			ipAddress := IPAddress{Address: ip, Type: tt.netType}

			require.NotNil(t, ipAddress.Address)
			require.Equal(t, tt.ip, ipAddress.Address.String())
			require.Equal(t, tt.netType, ipAddress.Type)
		})
	}

	tests = []struct {
		description string
		ip          string
		netType     string
	}{
		{
			description: "Test successful JSON serialization of IPAddress with valid IPv4 address",
			ip:          "192.168.1.1",
			netType:     "IPv4",
		},
		{
			description: "Test successful JSON serialization of IPAddress with valid IPv6 address",
			ip:          "2001:db8::1",
			netType:     "IPv6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			ip, _ := netip.ParseAddr(tt.ip)
			ipAddress := IPAddress{Address: ip, Type: tt.netType}

			jsonData, err := json.Marshal(ipAddress)

			require.NoError(t, err)
			require.JSONEq(t, fmt.Sprintf(`{"address":"%s","type":"%s"}`, tt.ip, tt.netType), string(jsonData))
		})
	}
}
