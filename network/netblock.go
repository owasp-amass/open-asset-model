package network

import (
	"encoding/json"
	"net/netip"

	openAssetModel "github.com/owasp-amass/open-asset-model"
)

// Netblock represents a block of IP addresses in a network.
// It is often used to specify a range of IP addresses that
// are assigned to a particular organization or network.
// Netblock implements Asset interface.
type Netblock struct {
	// Cidr is the CIDR notation of the IP address block,
	// such as "198.51.100.0/24" (IPv4) or "2001:db8::/32" (IPv6).
	Cidr netip.Prefix `json:"cidr"`

	// Type is the type of the IP address block,
	// such as "IPv4" or "IPv6".
	Type string `json:"type"`
}

// AssetType returns the asset type.
func (n Netblock) AssetType() openAssetModel.AssetType {
	return openAssetModel.Netblock
}

// JSON returns the JSON encoding of the struct.
func (n Netblock) JSON() ([]byte, error) {
	return json.Marshal(n)
}
