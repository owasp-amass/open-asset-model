package network_test

import (
	"fmt"
	"net/netip"
	"testing"

	openAssetModel "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/network"

	"github.com/stretchr/testify/require"
)

func TestNetblockImplementsAsset(t *testing.T) {
	var _ openAssetModel.Asset = Netblock{}       // Verify that Netblock implements Asset interface
	var _ openAssetModel.Asset = (*Netblock)(nil) // Verify that *Netblock implements Asset interface.
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
			netblock := Netblock{Cidr: prefix, Type: tt.netType}

			require.Equal(t, tt.cidr, netblock.Cidr.String())
			require.Equal(t, tt.netType, netblock.Type)

			require.Equal(t, openAssetModel.Netblock, netblock.AssetType())
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
			netblock := Netblock{Cidr: prefix, Type: tt.netType}

			jsonData, err := netblock.JSON()

			require.NoError(t, err)
			require.JSONEq(t, fmt.Sprintf(`{"cidr":"%s","type":"%s"}`, tt.cidr, tt.netType), string(jsonData))
		})
	}
}
