package org

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Organization represents an organization. Does not have to be tied with an RIR.
type Organization struct {
	// OrgName is the name of the organization.
	OrgName string `json:"org_name,omitempty"`
	// Industry is the industry of the organization.
	Industry string `json:"industry,omitempty"`
	// IANAid is the IANA ID of the organization.
	IANAid string `json:"iana_id,omitempty"`
}

// AssetType returns the asset type.
func (o Organization) AssetType() model.AssetType {
	return model.Organization
}

// JSON returns the JSON representation of the asset.
func (o Organization) JSON() ([]byte, error) {
	return json.Marshal(o)
}
