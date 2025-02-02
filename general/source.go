// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package general

import (
	"encoding/json"
	"strconv"

	model "github.com/owasp-amass/open-asset-model"
)

// SourceProperty represents a source of data in the graph.
type SourceProperty struct {
	Source     string `json:"name"`
	Confidence int    `json:"confidence"`
}

// Name implements the Property interface.
func (p SourceProperty) Name() string {
	return p.Source
}

// Value implements the Property interface.
func (p SourceProperty) Value() string {
	return strconv.Itoa(p.Confidence)
}

// PropertyType implements the Property interface.
func (p SourceProperty) PropertyType() model.PropertyType {
	return model.SourceProperty
}

// JSON implements the Property interface.
func (p SourceProperty) JSON() ([]byte, error) {
	return json.Marshal(p)
}
