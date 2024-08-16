// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// EmailAddress represents an email address with a value, local part, and domain.
type EmailAddress struct {
	// Address is the full email address.
	Address string `json:"address"`
	// LocalPart is the part of the email address before the "@" symbol.
	LocalPart string `json:"local_part"`
	// Domain is the part of the email address after the "@" symbol.
	Domain string `json:"domain"`
}

// AssetType returns the asset type.
func (e EmailAddress) AssetType() model.AssetType {
	return model.EmailAddress
}

// JSON returns the JSON representation of the asset.
func (e EmailAddress) JSON() ([]byte, error) {
	return json.Marshal(e)
}
