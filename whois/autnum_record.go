// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package whois

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// AutnumRecord represents the RDAP record for an autonomous system.
type AutnumRecord struct {
	Raw         string   `json:"raw,omitempty"`
	Number      int      `json:"number"`
	Handle      string   `json:"handle,omitempty"`
	Name        string   `json:"name,omitempty"`
	WhoisServer string   `json:"whois_server,omitempty"`
	CreatedDate string   `json:"created_date,omitempty"`
	UpdatedDate string   `json:"updated_date,omitempty"`
	Status      []string `json:"status,omitempty"`
}

// AssetType returns the asset type.
func (as AutnumRecord) AssetType() model.AssetType {
	return model.AutnumRecord
}

// JSON returns the JSON representation of the asset.
func (as AutnumRecord) JSON() ([]byte, error) {
	return json.Marshal(as)
}
