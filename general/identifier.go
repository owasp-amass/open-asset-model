// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package general

import (
	"encoding/json"
	"fmt"

	model "github.com/owasp-amass/open-asset-model"
)

// Known identifier types.
const (
	AccountNumber         = "account"    // Account number
	ASINNumber            = "asin"       // Amazon Standard Identification Number
	BankIDCode            = "bic"        // Bank Identifier Code
	DUNSNumer             = "duns"       // Data Universal Numbering System, Dun & Bradstreet (D&B) assigns this number
	EAN8                  = "ean_8"      // International Article Number (EAN) 8-digit
	EAN13                 = "ean_13"     // International Article Number (EAN) 13-digit
	EmailAddress          = "email"      // Email address
	EmployerIDNumber      = "ein"        // Employer Identification Number
	GlobalLocNumber       = "gln"        // Global Location Number
	GlobalTradeItemNumber = "gtin"       // Global Trade Item Number
	GovIssuedIDNumber     = "gov_id"     // Government issued identification number for individuals
	IBANCode              = "iban"       // International Bank Account Number
	ICANNAuthCode         = "icann_auth" // Auth Code, Internet Corporation for Assigned Names and Numbers
	ICANNEPPCode          = "icann_epp"  // Extensible Provisioning Protocol (EPP) status code, ICANN
	ISBN                  = "isbn"       // International Standard Book Number
	ISICCode              = "isic"       // International Standard Industrial Classification
	LEICode               = "lei"        // Legal Entity Identifier, Global Legal Entity Identifier Foundation assigns this number
	ModelNumber           = "model"      // Model number
	MPN                   = "mpn"        // Manufacturer Part Number
	NAICSCode             = "naics"      // North American Industry Classification System
	NSNCode               = "nsn"        // NATO Stock Number
	OECDCode              = "oecd"       // Organization for Economic Co-operation and Development
	SerialNumber          = "serial"     // Serial number
	SICCode               = "sic"        // Standard Industrial Classification
	SWIFTCode             = "swift"      // Society for Worldwide Interbank Financial Telecommunication
	TaxIDNumber           = "tin"        // Taxpayer Identification Number
	TickerSymbol          = "ticker"     // Stock ticker symbol
	UPCNumberA            = "upc_a"      // Universal Product Code (UPC) Version A
	UPCNumberE            = "upc_e"      // Universal Product Code (UPC) Version E
	VersionNumber         = "version"    // Version number
)

// Identifier identifies something that's a member of a system or organization that issues ID numbers or codes.
// Should support relationships for the following:
// - Registration agency (e.g., ContactRecord)
// - Issuing authority (e.g., ContactRecord)
// - Issuing agent (e.g., ContactRecord)
type Identifier struct {
	ID             string `json:"id"`
	Type           string `json:"id_type"`
	Category       string `json:"category,omitempty"`
	CreationDate   string `json:"creation_date,omitempty"`
	UpdatedDate    string `json:"update_date,omitempty"`
	ExpirationDate string `json:"expiration_date,omitempty"`
	Status         string `json:"status,omitempty"`
}

// Key implements the Asset interface.
func (r Identifier) Key() string {
	return fmt.Sprintf("%s:%s", r.Type, r.ID)
}

// AssetType implements the Asset interface.
func (r Identifier) AssetType() model.AssetType {
	return model.Identifier
}

// JSON implements the Asset interface.
func (r Identifier) JSON() ([]byte, error) {
	return json.Marshal(r)
}
