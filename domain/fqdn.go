// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package domain

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// FQDN represents a Fully Qualified Domain Name.
type FQDN struct {
	// Name is the domain name part of the FQDN.
	// It should be a valid domain name, such as "example" or "subdomain.example.com".
	Name string `json:"name"`
}

// AssetType returns the asset type.
func (f FQDN) AssetType() model.AssetType {
	return model.FQDN
}

// JSON returns the JSON encoding of the struct.
func (f FQDN) JSON() ([]byte, error) {
	return json.Marshal(f)
}
