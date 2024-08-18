// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// EmailAddress represents an email address with an address, username, and the domain.
type EmailAddress struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Domain   string `json:"domain"`
}

// AssetType returns the asset type.
func (e EmailAddress) AssetType() model.AssetType {
	return model.EmailAddress
}

// JSON returns the JSON representation of the asset.
func (e EmailAddress) JSON() ([]byte, error) {
	return json.Marshal(e)
}
