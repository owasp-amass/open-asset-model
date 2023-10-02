package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// This struct represents the phone number, whether it is fax, mobile, or home number linked to the possible asset
type Phone struct {
	Type  string `json:"type"`
	Phone string `json:"phone"`
	Ext   string `json:"ext,omitempty"`
}

// AssetType returns the asset type.
func (p Phone) AssetType() model.AssetType {
	return model.Phone
}

// JSON returns the JSON representation of the asset.
func (p Phone) JSON() ([]byte, error) {
	return json.Marshal(p)
}
