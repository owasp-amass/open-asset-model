// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// This struct represents the location linked to the possible asset
type Location struct {
	// FormattedAddress is the human-readable address of the location
	FormattedAddress string `json:"formatted_address,omitempty"`
	BuildingNumber   string `json:"building_number,omitempty"`
	StreetName       string `json:"street_name,omitempty"`
	Unit             string `json:"unit,omitempty"`
	Building         string `json:"building,omitempty"`
	City             string `json:"city,omitempty"`
	Locality         string `json:"locality,omitempty"`
	Province         string `json:"province,omitempty"`
	CountryCode      string `json:"country_code,omitempty"` // ISO 3166-1 alpha-2 country code
	PostalCode       string `json:"postal_code,omitempty"`
}

// AssetType returns the asset type.
func (a Location) AssetType() model.AssetType {
	return model.Location
}

// JSON returns the JSON representation of the asset.
func (a Location) JSON() ([]byte, error) {
	return json.Marshal(a)
}
