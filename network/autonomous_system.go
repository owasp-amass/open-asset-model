// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"encoding/json"
	"strconv"

	model "github.com/owasp-amass/open-asset-model"
)

// AutonomousSystem represents an autonomous system.
type AutonomousSystem struct {
	Number int `json:"number"`
}

// Key implements the Asset interface.
func (a AutonomousSystem) Key() string {
	return strconv.Itoa(a.Number)
}

// AssetType implements the Asset interface.
func (a AutonomousSystem) AssetType() model.AssetType {
	return model.AutonomousSystem
}

// JSON implements the Asset interface.
func (a AutonomousSystem) JSON() ([]byte, error) {
	return json.Marshal(a)
}
