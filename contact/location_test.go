// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestLocationAssetType(t *testing.T) {
	loc := Location{}
	if loc.AssetType() != model.Location {
		t.Errorf("Expected asset type %s but got %s", model.Location, loc.AssetType())
	}
}

func TestLocationJSON(t *testing.T) {
	loc := Location{
		FormattedAddress: "123 Main St",
		BuildingNumber:   "123",
		StreetName:       "Main St",
		Unit:             "Apt 1",
		Building:         "Building A",
		City:             "Anytown",
		Locality:         "Anytown",
		Province:         "Anyregion",
		CountryCode:      "US",
		PostalCode:       "12345",
	}

	expectedJSON := `{"formatted_address":"123 Main St","building_number":"123","street_name":"Main St","unit":"Apt 1","building":"Building A","city":"Anytown","locality":"Anytown","province":"Anyregion","country_code":"US","postal_code":"12345"}`

	jsonData, err := loc.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s but got %s", expectedJSON, string(jsonData))
	}
}
