// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package property

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// VulnProperty represents a simple property in the graph.
type VulnProperty struct {
	ID          string `json:"id"`
	Description string `json:"desc"`
	Source      string `json:"source"`
	Category    string `json:"category"`
	Enumeration string `json:"enum"`
	Reference   string `json:"ref"`
}

// Name implements the Property interface.
func (p VulnProperty) Name() string {
	return p.ID
}

// Value implements the Property interface.
func (p VulnProperty) Value() string {
	return p.Description
}

// PropertyType implements the Property interface.
func (p VulnProperty) PropertyType() model.PropertyType {
	return model.VulnProperty
}

// JSON implements the Property interface.
func (p VulnProperty) JSON() ([]byte, error) {
	return json.Marshal(p)
}
