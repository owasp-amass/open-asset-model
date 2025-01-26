// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package org

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Organization represents an organization.
// Should support relationships for the following:
// - Principal place of business
// - Parent organizations
// - Subsidiaries
// - Sister companies
// - DUNS number
// - Tax identification number
// - Trader identification number
// - ARIN organization identifier
// - Decentralized identifier (DID)
// - Ticker symbol
// - Global Location Number (GLN)
// - ISIC, NAICS, SIC, BIC, and ISO 6523 code
// - Legal Entity Identifier (LEI) ISO 17442 code
// - Registration number
// - Website
// - Social media profiles
// - Contact information
// - Founder, sponsorships, and funding sources
type Organization struct {
	Name           string `json:"name"`
	LegalName      string `json:"legal_name,omitempty"`
	FoundingDate   string `json:"founding_date,omitempty"`
	Industry       string `json:"industry,omitempty"`
	Active         bool   `json:"active,omitempty"`
	NonProfit      bool   `json:"non_profit,omitempty"`
	NumOfEmployees int    `json:"num_of_employees,omitempty"`
}

// Key implements the Asset interface.
func (o Organization) Key() string {
	return o.Name
}

// AssetType implements the Asset interface.
func (o Organization) AssetType() model.AssetType {
	return model.Organization
}

// JSON implements the Asset interface.
func (o Organization) JSON() ([]byte, error) {
	return json.Marshal(o)
}
