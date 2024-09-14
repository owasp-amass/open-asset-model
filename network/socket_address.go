// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"encoding/json"
	"net/netip"

	model "github.com/owasp-amass/open-asset-model"
)

type SocketAddress struct {
	Address netip.AddrPort `json:"address"`

	// IPAddress is the IP address value, stored as a netip.Addr type.
	// It should be a valid IP address, such as "192.0.2.1" (IPv4)
	// or "2001:db8::1" (IPv6).
	IPAddress netip.Addr `json:"ip_address"`

	// Port is the port number.
	Port int `json:"port"`

	// Protocol is the protocol used to reach the port, such as "tcp" or "udp".
	Protocol string `json:"protocol"`
}

// Key implements the Asset interface.
func (sa SocketAddress) Key() string {
	return sa.Address.String()
}

// AssetType implements the Asset interface.
func (sa SocketAddress) AssetType() model.AssetType {
	return model.SocketAddress
}

// JSON implements the Asset interface.
func (sa SocketAddress) JSON() ([]byte, error) {
	return json.Marshal(sa)
}
