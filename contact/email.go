package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Email address linked to an asset
type Email struct {
	Email string `json:"email"`
}

// AssetType returns the asset type.
func (e Email) AssetType() model.AssetType {
	return model.Email
}

// JSON returns the JSON representation of the asset.
func (e Email) JSON() ([]byte, error) {
	return json.Marshal(e)
}
