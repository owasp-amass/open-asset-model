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

func TestNetblockKey(t *testing.T) {
	want := "192.168.1.0/24"
	nb := Netblock{CIDR: netip.MustParsePrefix(want)}

	if got := nb.Key(); got != want {
		t.Errorf("Netblock.Key() = %v, want %v", got, want)
	}
}

func TestNetblockImplementsAsset(t *testing.T) {
	var _ model.Asset = Netblock{}       // Verify that Netblock implements Asset interface
	var _ model.Asset = (*Netblock)(nil) // Verify that *Netblock implements Asset interface.
}

func TestNetblock(t *testing.T) {
	tests := []struct {
		description string
		cidr        string
		netType     string
	}{
		{
			description: "Test successful creation of Netblock with valid IPv4 CIDR",
			cidr:        "198.51.100.0/24",
			netType:     "IPv4",
		},
		{
			description: "Test successful creation of Netblock with valid IPv6 CIDR",
			cidr:        "2001:db8::/32",
			netType:     "IPv6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			prefix, _ := netip.ParsePrefix(tt.cidr)
			netblock := Netblock{CIDR: prefix, Type: tt.netType}

			require.Equal(t, tt.cidr, netblock.CIDR.String())
			require.Equal(t, tt.netType, netblock.Type)

			require.Equal(t, model.Netblock, netblock.AssetType())
		})
	}

	tests = []struct {
		description string
		cidr        string
		netType     string
	}{
		{
			description: "Test successful JSON serialization of Netblock with valid IPv4 CIDR",
			cidr:        "198.51.100.0/24",
			netType:     "IPv4",
		},
		{
			description: "Test successful JSON serialization of Netblock with valid IPv6 CIDR",
			cidr:        "2001:db8::/32",
			netType:     "IPv6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			prefix, _ := netip.ParsePrefix(tt.cidr)
			netblock := Netblock{CIDR: prefix, Type: tt.netType}

			jsonData, err := netblock.JSON()

			require.NoError(t, err)
			require.JSONEq(t, fmt.Sprintf(`{"cidr":"%s","type":"%s"}`, tt.cidr, tt.netType), string(jsonData))
		})
	}
}
