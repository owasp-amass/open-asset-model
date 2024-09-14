// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package service

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

type Service struct {
	Identifier string `json:"identifier"`
	Version    string `json:"version"`
}

// Key implements the Asset interface.
func (s Service) Key() string {
	return s.Identifier
}

// AssetType implements the Asset interface.
func (s Service) AssetType() model.AssetType {
	return model.Service
}

// JSON implements the Asset interface.
func (s Service) JSON() ([]byte, error) {
	return json.Marshal(s)
}
