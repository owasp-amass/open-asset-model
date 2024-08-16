// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package people

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Person represents a person's information.
type Person struct {
	FullName   string `json:"full_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name,omitempty"`
	FamilyName string `json:"family_name"`
}

// AssetType returns the asset type.
func (n Person) AssetType() model.AssetType {
	return model.Person
}

// JSON returns the JSON representation of the asset.
func (n Person) JSON() ([]byte, error) {
	return json.Marshal(n)
}
