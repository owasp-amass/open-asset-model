// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package financial

import (
	"encoding/json"
	"fmt"

	model "github.com/owasp-amass/open-asset-model"
)

// Account represents a financial account managed by a financial institution.
// Should support relationships for the following:
// - Financial institution (e.g. Organization)
// - Owners (e.g. Person or Organization)
// - Funds transfers
// - IBAN and SWIFT codes
type Account struct {
	Number             string  `json:"account_number"`
	Type               string  `json:"account_type"`
	DomesticBankNumber string  `json:"domestic_bank_number,omitempty"`
	Balance            int     `json:"balance,omitempty"`
	Currency           string  `json:"currency,omitempty"`
	InterestRate       float64 `json:"interest_rate,omitempty"`
	PercentageRate     float64 `json:"percentage_rate,omitempty"`
	FeesAndCommissions int     `json:"fees_and_commissions,omitempty"`
	Active             bool    `json:"active,omitempty"`
	CountryCode        string  `json:"country_code,omitempty"`
}

// Key implements the Asset interface.
func (a Account) Key() string {
	return fmt.Sprintf("%s:%s", a.Type, a.Number)
}

// AssetType implements the Asset interface.
func (a Account) AssetType() model.AssetType {
	return model.Account
}

// JSON implements the Asset interface.
func (a Account) JSON() ([]byte, error) {
	return json.Marshal(a)
}
