// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

const (
	PhoneTypeRegular string = "phone"
	PhoneTypeFax     string = "fax"
	PhoneTypeMobile  string = "mobile"
)

// This struct represents the phone number, whether it is fax, mobile, or home number linked to the possible asset
type Phone struct {
	Type          string `json:"type,omitempty"`
	Raw           string `json:"raw"`
	E164          string `json:"e164"` // E.164 format
	CountryAbbrev string `json:"country_abbrev,omitempty"`
	CountryCode   int    `json:"country_code,omitempty"`
	Ext           string `json:"ext,omitempty"`
}

// AssetType returns the asset type.
func (p Phone) AssetType() model.AssetType {
	return model.Phone
}

// JSON returns the JSON representation of the asset.
func (p Phone) JSON() ([]byte, error) {
	return json.Marshal(p)
}
