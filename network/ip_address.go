// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"encoding/json"
	"net/netip"

	model "github.com/owasp-amass/open-asset-model"
)

// IPAddress represents an IP address.
type IPAddress struct {
	// Address is the IP address value, stored as a netip.Addr type.
	// It should be a valid IP address, such as "192.0.2.1" (IPv4)
	// or "2001:db8::1" (IPv6).
	Address netip.Addr `json:"address"`

	// Type is the type of IP address, such as "IPv4" or "IPv6".
	Type string `json:"type"`
}

// Key implements the Asset interface.
func (ip IPAddress) Key() string {
	return ip.Address.String()
}

// AssetType implements the Asset interface.
func (ip IPAddress) AssetType() model.AssetType {
	return model.IPAddress
}

// JSON implements the Asset interface.
func (ip IPAddress) JSON() ([]byte, error) {
	return json.Marshal(ip)
}
