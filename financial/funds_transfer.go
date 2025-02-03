// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package financial

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// FundsTransfer represents a single transfer of funds between two accounts.
// Should support relationships for the following:
// - Sender financial account (e.g. Account)
// - Recipient financial account (e.g. Account)
// - IBIN and SWIFT codes
type FundsTransfer struct {
	ID              string  `json:"unique_id"`
	Amount          float64 `json:"amount"`
	ReferenceNumber string  `json:"reference_number,omitempty"`
	Currency        string  `json:"currency,omitempty"`
	Method          string  `json:"transfer_method,omitempty"`
	ExchangeDate    string  `json:"exchange_date,omitempty"`
	ExchangeRate    float64 `json:"exchange_rate,omitempty"`
}

// Key implements the Asset interface.
func (ft FundsTransfer) Key() string {
	return ft.ID
}

// AssetType implements the Asset interface.
func (ft FundsTransfer) AssetType() model.AssetType {
	return model.FundsTransfer
}

// JSON implements the Asset interface.
func (ft FundsTransfer) JSON() ([]byte, error) {
	return json.Marshal(ft)
}
