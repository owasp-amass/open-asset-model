// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package general

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// SimpleRelation represents a simple relation in the graph with no additional data required.
type SimpleRelation struct {
	Name string `json:"label"`
}

// Label implements the Relation interface.
func (r SimpleRelation) Label() string {
	return r.Name
}

// RelationType implements the Relation interface.
func (r SimpleRelation) RelationType() model.RelationType {
	return model.SimpleRelation
}

// JSON implements the Relation interface.
func (r SimpleRelation) JSON() ([]byte, error) {
	return json.Marshal(r)
}

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
