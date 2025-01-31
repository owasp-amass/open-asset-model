// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"encoding/json"
	"net/netip"

	model "github.com/owasp-amass/open-asset-model"
)

// Netblock represents a block of IP addresses in a network.
// It is often used to specify a range of IP addresses that
// are assigned to a particular organization or network.
type Netblock struct {
	// CIDR is the Classless Inter-Domain Routing notation of
	// the IP address block, such as "198.51.100.0/24" (IPv4)
	// or "2001:db8::/32" (IPv6)
	CIDR netip.Prefix `json:"cidr"`

	// Type is the type of the IP address block,
	// such as "IPv4" or "IPv6"
	Type string `json:"type"`
}

// Key implements the Asset interface.
func (nb Netblock) Key() string {
	return nb.CIDR.String()
}

// AssetType implements the Asset interface.
func (nb Netblock) AssetType() model.AssetType {
	return model.Netblock
}

// JSON implements the Asset interface.
func (nb Netblock) JSON() ([]byte, error) {
	return json.Marshal(nb)
}
