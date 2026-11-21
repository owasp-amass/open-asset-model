// Copyright © by Jeff Foley 2017-2026. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// StreetAddress represents a street or postal address, as written, with its
// parsed parts. Its identity is the full address string. A position on the
// Earth is a GeoLocation, linked from the address when one is known
// (StreetAddress --geolocation--> GeoLocation).
type StreetAddress struct {
	Address        string `json:"address"`
	Building       string `json:"building,omitempty"`
	BuildingNumber string `json:"building_number,omitempty"`
	StreetName     string `json:"street_name,omitempty"`
	Unit           string `json:"unit,omitempty"`
	POBox          string `json:"po_box,omitempty"`
	City           string `json:"city"`
	Locality       string `json:"locality,omitempty"`
	Province       string `json:"province,omitempty"`
	Country        string `json:"country,omitempty"`
	PostalCode     string `json:"postal_code,omitempty"`
	GLN            int    `json:"gln,omitempty"`
}

// Key implements the Asset interface.
func (a StreetAddress) Key() string {
	return a.Address
}

// AssetType implements the Asset interface.
func (a StreetAddress) AssetType() model.AssetType {
	return model.StreetAddress
}

// JSON implements the Asset interface.
func (a StreetAddress) JSON() ([]byte, error) {
	return json.Marshal(a)
}
