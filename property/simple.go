// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package property

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// SimpleProperty represents a simple property in the graph.
type SimpleProperty struct {
	PropertyName  string `json:"property_name"`
	PropertyValue string `json:"property_value"`
}

// Name implements the Property interface.
func (p SimpleProperty) Name() string {
	return p.PropertyName
}

// Value implements the Property interface.
func (p SimpleProperty) Value() string {
	return p.PropertyValue
}

// PropertyType implements the Property interface.
func (p SimpleProperty) PropertyType() model.PropertyType {
	return model.SimpleProperty
}

// JSON implements the Property interface.
func (p SimpleProperty) JSON() ([]byte, error) {
	return json.Marshal(p)
}
