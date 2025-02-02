// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// FQDN represents a Fully Qualified Domain Name.
type FQDN struct {
	// it should be a valid domain name, such as "subdomain.example.com"
	Name string `json:"name"`
}

// Key implements the Asset interface.
func (f FQDN) Key() string {
	return f.Name
}

// AssetType implements the Asset interface.
func (f FQDN) AssetType() model.AssetType {
	return model.FQDN
}

// JSON implements the Asset interface.
func (f FQDN) JSON() ([]byte, error) {
	return json.Marshal(f)
}
