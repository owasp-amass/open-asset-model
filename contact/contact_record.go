// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// This type links various assets together to form complete contact information
type ContactRecord struct {
	DiscoveredAt string `json:"discovered_at,omitempty"`
}

// Key implements the Asset interface.
func (cr ContactRecord) Key() string {
	return cr.DiscoveredAt
}

// AssetType implements the Asset interface.
func (cr ContactRecord) AssetType() model.AssetType {
	return model.ContactRecord
}

// JSON implements the Asset interface.
func (cr ContactRecord) JSON() ([]byte, error) {
	return json.Marshal(cr)
}
