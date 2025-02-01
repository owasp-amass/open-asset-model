// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package account

import (
	"encoding/json"
	"fmt"

	model "github.com/owasp-amass/open-asset-model"
)

// Account represents an account managed by an organization.
// Should support relationships for the following:
// - User (e.g. Person or Organization)
// - Funds transfers
// - IBAN and SWIFT codes
type Account struct {
	Username string `json:"username"`
	Number   string `json:"account_number"`
	Type     string `json:"account_type"`
	Balance  int    `json:"balance,omitempty"`
	Active   bool   `json:"active,omitempty"`
}

// Key implements the Asset interface.
func (a Account) Key() string {
	return fmt.Sprintf("%s:%s:%s", a.Type, a.Username, a.Number)
}

// AssetType implements the Asset interface.
func (a Account) AssetType() model.AssetType {
	return model.Account
}

// JSON implements the Asset interface.
func (a Account) JSON() ([]byte, error) {
	return json.Marshal(a)
}
