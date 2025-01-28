// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package platform

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Service represents a service provided by an asset and/or organization.
// It should support relationships such as the following:
// - Providers
// - Products used
// - Terms of service
type Service struct {
	UniqueID   string              `json:"unique_id"`
	Type       string              `json:"service_type"`
	Output     string              `json:"output"`
	OutputLen  int                 `json:"output_length"`
	Attributes map[string][]string `json:"attributes,omitempty"`
}

// Key implements the Asset interface.
func (s Service) Key() string {
	return s.UniqueID
}

// AssetType implements the Asset interface.
func (s Service) AssetType() model.AssetType {
	return model.Service
}

// JSON implements the Asset interface.
func (s Service) JSON() ([]byte, error) {
	return json.Marshal(s)
}
