// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// EmailAddress represents an email address with an address, username, and the domain.
type EmailAddress struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Domain   string `json:"domain"`
}

// Key implements the Asset interface.
func (e EmailAddress) Key() string {
	return e.Address
}

// AssetType implements the Asset interface.
func (e EmailAddress) AssetType() model.AssetType {
	return model.EmailAddress
}

// JSON implements the Asset interface.
func (e EmailAddress) JSON() ([]byte, error) {
	return json.Marshal(e)
}
