package people

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Name of a person.
type Name struct {
	Name string `json:"name"`
}

// AssetType returns the asset type.
func (n Name) AssetType() model.AssetType {
	return model.Name
}

// JSON returns the JSON representation of the asset.
func (n Name) JSON() ([]byte, error) {
	return json.Marshal(n)
}
