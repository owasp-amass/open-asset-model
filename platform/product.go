// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package platform

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Product represents a technology product and organizes the product releases in the data model.
// Should support relationships for the following:
// - Manufacturer (e.g. Organization)
// - Website
// - Product releases
type Product struct {
	ID              string `json:"unique_id"`
	Name            string `json:"product_name"`
	Type            string `json:"product_type"`
	Category        string `json:"category,omitempty"`
	Description     string `json:"description,omitempty"`
	CountryOfOrigin string `json:"country_of_origin,omitempty"`
}

// Key implements the Asset interface.
func (p Product) Key() string {
	return p.ID
}

// AssetType implements the Asset interface.
func (p Product) AssetType() model.AssetType {
	return model.Product
}

// JSON implements the Asset interface.
func (p Product) JSON() ([]byte, error) {
	return json.Marshal(p)
}

// ProductRelease represents a release of a technology product that belongs to a Product.
// Should support relationships for the following:
// - Amazon Standard Identification Number (ASIN)
// - Global Trade Item Number (GTIN)
// - International Article Number (EAN)
// - International Standard Book Number (ISBN)
// - Manufacturer Part Number (MPN)
// - Model Number
// - NATO Stock Number (NSN)
// - Serial Number
// - Universal Product Code (UPC)
// - Version Number
// - Vulnerabilities
// - Website with release details
type ProductRelease struct {
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date,omitempty"`
}

// Key implements the Asset interface.
func (p ProductRelease) Key() string {
	return p.Name
}

// AssetType implements the Asset interface.
func (p ProductRelease) AssetType() model.AssetType {
	return model.ProductRelease
}

// JSON implements the Asset interface.
func (p ProductRelease) JSON() ([]byte, error) {
	return json.Marshal(p)
}
