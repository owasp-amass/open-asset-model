// Copyright © by Jeff Foley 2017-2026. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package file

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// File represents a file discovered, such as a document or image. Its
// identity is the URL it was retrieved from; the same content found at several
// URLs is linked through the File "fingerprint" relation to a shared
// Identifier holding the content hash. Metadata read from the file (author,
// software, dates) is recorded as properties, and the people, email addresses
// and locations it names through the "contains" relation to a ContactRecord.
type File struct {
	URL  string `json:"url"`
	Name string `json:"name,omitempty"`
	// Type is the file's media type (RFC 6838), e.g. "application/pdf" or
	// "image/jpeg".
	Type string `json:"type,omitempty"`
	// Size is the file's length in bytes, when known.
	Size int64 `json:"size,omitempty"`
}

// Key implements the Asset interface.
func (f File) Key() string {
	return f.URL
}

// AssetType implements the Asset interface.
func (f File) AssetType() model.AssetType {
	return model.File
}

// JSON implements the Asset interface.
func (f File) JSON() ([]byte, error) {
	return json.Marshal(f)
}
