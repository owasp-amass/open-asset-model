// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package id

import (
	"encoding/json"
	"fmt"

	model "github.com/owasp-amass/open-asset-model"
)

// Identifier identifies something that's a member of a system or organization that issues ID numbers or codes.
// Should support relationships for the following:
// - Member (e.g., person, organization, or device)
// - Registration location
// - Registration authority
// - Issuing authority
// - Issuing authority Website
// - Agent of authority (e.g., person or organization)
type Identifier struct {
	ID             string `json:"id"`
	Type           string `json:"id_type"`
	Category       string `json:"category,omitempty"`
	CreationDate   string `json:"creation_date,omitempty"`
	UpdatedDate    string `json:"update_date,omitempty"`
	ExpirationDate string `json:"expiration_date,omitempty"`
	Status         string `json:"status,omitempty"`
}

// Key implements the Asset interface.
func (r Identifier) Key() string {
	return fmt.Sprintf("%s:%s", r.Type, r.ID)
}

// AssetType implements the Asset interface.
func (r Identifier) AssetType() model.AssetType {
	return model.Identifier
}

// JSON implements the Asset interface.
func (r Identifier) JSON() ([]byte, error) {
	return json.Marshal(r)
}
