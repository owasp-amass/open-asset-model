// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package whois

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Registrar represents a domain registrar.
type Registrar struct {
	// Name is the name of the registrar.
	Name string `json:"name"`
	// URL is the URL of the registrar's website.
	URL string `json:"url"`
	// IANAID is the IANA ID of the registrar.
	IANAID string `json:"iana_id"`
}

func (r Registrar) AssetType() model.AssetType {
	return model.Registrar
}

func (r Registrar) JSON() ([]byte, error) {
	return json.Marshal(r)
}
