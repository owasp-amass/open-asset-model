package network

import (
	"encoding/json"
	"net/netip"

	openAssetModel "github.com/owasp-amass/open-asset-model"
)

// IPAddress represents an IP address.
// IPAddress implements Asset interface.
type IPAddress struct {
	// Address is the IP address value, stored as a netip.Addr type.
	// It should be a valid IP address, such as "192.0.2.1" (IPv4)
	// or "2001:db8::1" (IPv6).
	Address netip.Addr `json:"address"`

	// Type is the type of IP address, such as "IPv4" or "IPv6".
	Type string `json:"type"`
}

// AssetType returns the asset type.
func (i IPAddress) AssetType() openAssetModel.AssetType {
	return openAssetModel.IPAddress
}

// JSON returns the JSON encoding of the struct.
func (i IPAddress) JSON() ([]byte, error) {
	return json.Marshal(i)
}
