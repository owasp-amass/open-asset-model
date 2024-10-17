// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package file

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// File represents a file discovered, such as a document or image.
type File struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// Key implements the Asset interface.
func (f File) Key() string {
	return f.Name
}

// AssetType implements the Asset interface.
func (f File) AssetType() model.AssetType {
	return model.File
}

// JSON implements the Asset interface.
func (f File) JSON() ([]byte, error) {
	return json.Marshal(f)
}
