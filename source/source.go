// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package source

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Source represents a data source for a piece of open source intelligence.
type Source struct {
	// Name is the unique name given to the data source.
	Name string `json:"name"`
	// Confidence represents the percentage (0-100) of trust in the source.
	Confidence int `json:"confidence"`
}

// AssetType returns the asset type.
func (s Source) AssetType() model.AssetType {
	return model.Source
}

// JSON returns the JSON encoding of the struct.
func (s Source) JSON() ([]byte, error) {
	return json.Marshal(s)
}
