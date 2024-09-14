// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"fmt"
	"net/netip"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestIPAddressKey(t *testing.T) {
	want := "1.1.1.1"
	ip := IPAddress{Address: netip.MustParseAddr(want)}

	if got := ip.Key(); got != want {
		t.Errorf("IPAddress.Key() = %v, want %v", got, want)
	}
}

func TestIPAddressAsset(t *testing.T) {
	var _ model.Asset = IPAddress{}       // Verify that IPAddress implements Asset interface
	var _ model.Asset = (*IPAddress)(nil) // Verify that *IPAddress implements Asset interface.
}

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
			require.Equal(t, ipAddress.AssetType(), model.IPAddress)
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

			jsonData, err := ipAddress.JSON()

			require.NoError(t, err)
			require.JSONEq(t, fmt.Sprintf(`{"address":"%s","type":"%s"}`, tt.ip, tt.netType), string(jsonData))
		})
	}
}
