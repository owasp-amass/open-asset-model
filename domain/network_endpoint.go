// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package domain

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

type NetworkEndpoint struct {
	Address string `json:"address"`

	// Name is the domain name portion of the address
	Name string `json:"name"`

	// Port is the port number portion of the address
	Port int `json:"port"`

	// Protocol is the protocol used to reach the port, such as "tcp" or "udp".
	Protocol string `json:"protocol"`
}

// Key implements the Asset interface.
func (ne NetworkEndpoint) Key() string {
	return ne.Address
}

// AssetType implements the Asset interface.
func (ne NetworkEndpoint) AssetType() model.AssetType {
	return model.NetworkEndpoint
}

// JSON implements the Asset interface.
func (ne NetworkEndpoint) JSON() ([]byte, error) {
	return json.Marshal(ne)
}
