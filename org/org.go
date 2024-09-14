// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package org

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Organization represents an organization.
type Organization struct {
	Name     string `json:"name,omitempty"`
	Industry string `json:"industry,omitempty"`
}

// Key implements the Asset interface.
func (o Organization) Key() string {
	return o.Name
}

// AssetType implements the Asset interface.
func (o Organization) AssetType() model.AssetType {
	return model.Organization
}

// JSON implements the Asset interface.
func (o Organization) JSON() ([]byte, error) {
	return json.Marshal(o)
}
