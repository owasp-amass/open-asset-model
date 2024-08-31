// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package url

import (
	"bytes"
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// URL represents a URL.
type URL struct {
	Raw      string `json:"url"`                // Raw, unprocessed URL
	Scheme   string `json:"scheme"`             // Scheme (http, https)
	Username string `json:"username,omitempty"` // Username for authentication
	Password string `json:"password,omitempty"` // Password for authentication
	Host     string `json:"host"`               // Host
	Port     int    `json:"port,omitempty"`     // Port
	Path     string `json:"path"`               // Name
	Options  string `json:"options,omitempty"`  // Extra options used while connecting
	Fragment string `json:"fragment,omitempty"` // Fragment used in URI
}

// AssetType returns the asset type.
func (u URL) AssetType() model.AssetType {
	return model.URL
}

// JSON returns the JSON encoding of the struct.
func (u URL) JSON() ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(u)
	if err != nil {
		return nil, err
	}

	// Trim the newline added by the encoder
	return bytes.TrimRight(buffer.Bytes(), "\n"), nil
}
