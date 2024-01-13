package network

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

type Port struct {
	// Number is the port number.
	Number int `json:"number"`
	// Protocol is the protocol of the port, such as "tcp" or "udp".
	Protocol string `json:"protocol"`
}

// AssetType returns the asset type.
func (p Port) AssetType() model.AssetType {
	return model.Port
}

// JSON returns the JSON encoding of the struct.
func (p Port) JSON() ([]byte, error) {
	return json.Marshal(p)
}
