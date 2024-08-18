// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package org

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Organization represents an organization. Does not have to be tied with an RIR.
type Organization struct {
	Name     string `json:"name,omitempty"`
	Industry string `json:"industry,omitempty"`
}

// AssetType returns the asset type.
func (o Organization) AssetType() model.AssetType {
	return model.Organization
}

// JSON returns the JSON representation of the asset.
func (o Organization) JSON() ([]byte, error) {
	return json.Marshal(o)
}
