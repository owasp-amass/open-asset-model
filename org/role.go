package org

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Role of a corresponding person or associated with another asset (e.g. WHOIS records roles)
type Role struct {
	Role string `json:"role"`
}

// AssetType returns the asset type.
func (r Role) AssetType() model.AssetType {
	return model.Role
}

// JSON returns the JSON representation of the asset.
func (r Role) JSON() ([]byte, error) {
	return json.Marshal(r)
}
