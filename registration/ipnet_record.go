// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package registration

import (
	"encoding/json"
	"net/netip"

	model "github.com/owasp-amass/open-asset-model"
)

// IPNetRecord represents the RDAP record for an IP network.
type IPNetRecord struct {
	Raw          string       `json:"raw,omitempty"`
	CIDR         netip.Prefix `json:"cidr"`
	Handle       string       `json:"handle"`
	StartAddress netip.Addr   `json:"start_address"`
	EndAddress   netip.Addr   `json:"end_address"`
	Type         string       `json:"type"`
	Name         string       `json:"name"`
	Method       string       `json:"method,omitempty"`
	Country      string       `json:"country,omitempty"`
	ParentHandle string       `json:"parent_handle,omitempty"`
	WhoisServer  string       `json:"whois_server,omitempty"`
	CreatedDate  string       `json:"created_date"`
	UpdatedDate  string       `json:"updated_date"`
	Status       []string     `json:"status,omitempty"`
}

// Key implements the Asset interface.
func (ip IPNetRecord) Key() string {
	return ip.Handle
}

// AssetType implements the Asset interface.
func (ip IPNetRecord) AssetType() model.AssetType {
	return model.IPNetRecord
}

// JSON implements the Asset interface.
func (ip IPNetRecord) JSON() ([]byte, error) {
	return json.Marshal(ip)
}
