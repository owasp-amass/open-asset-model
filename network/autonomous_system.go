// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// AutonomousSystem represents an autonomous system.
type AutonomousSystem struct {
	// Number is the autonomous system number (AS).
	Number int `json:"number"`
}

// AssetType returns the asset type.
func (a AutonomousSystem) AssetType() model.AssetType {
	return model.ASN
}

// JSON returns the JSON encoding of the struct.
func (a AutonomousSystem) JSON() ([]byte, error) {
	return json.Marshal(a)
}
