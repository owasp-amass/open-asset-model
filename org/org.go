package org

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Organization represents an organization. Does not have to be tied with an RIR.
type Organization struct {
	Organization string `json:"organization"`
}

// AssetType returns the asset type.
func (o Organization) AssetType() model.AssetType {
	return model.Organization
}

// JSON returns the JSON representation of the asset.
func (o Organization) JSON() ([]byte, error) {
	return json.Marshal(o)
}

// The reason why I included this is because rir_org Refers to organizations that are tied with RIR.
// Not all organizations are tied to an R I R. For example, the organization that owns the domain name.
