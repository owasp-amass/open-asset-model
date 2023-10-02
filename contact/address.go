package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// This struct represents the address linked to the possible asset
type Address struct {
	Street        string `json:"street,omitempty"`
	City          string `json:"city,omitempty"`
	StateProvince string `json:"state/province,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	Country       string `json:"country,omitempty"`
}

// AssetType returns the asset type.
func (a Address) AssetType() model.AssetType {
	return model.Address
}

// JSON returns the JSON representation of the asset.
func (a Address) JSON() ([]byte, error) {
	return json.Marshal(a)
}
