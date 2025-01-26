// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package registration

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// DomainRecord represents the WHOIS record of a domain.
type DomainRecord struct {
	Raw            string   `json:"raw,omitempty"`
	ID             string   `json:"id,omitempty"`
	Domain         string   `json:"domain,omitempty"`
	Punycode       string   `json:"punycode,omitempty"`
	Name           string   `json:"name,omitempty"`
	Extension      string   `json:"extension,omitempty"`
	WhoisServer    string   `json:"whois_server,omitempty"`
	CreatedDate    string   `json:"created_date,omitempty"`
	UpdatedDate    string   `json:"updated_date,omitempty"`
	ExpirationDate string   `json:"expiration_date,omitempty"`
	Status         []string `json:"status,omitempty"`
	DNSSEC         bool     `json:"dnssec,omitempty"`
}

// Key implements the Asset interface.
func (dr DomainRecord) Key() string {
	return dr.Domain
}

// AssetType implements the Asset interface.
func (dr DomainRecord) AssetType() model.AssetType {
	return model.DomainRecord
}

// JSON implements the Asset interface.
func (dr DomainRecord) JSON() ([]byte, error) {
	return json.Marshal(dr)
}
