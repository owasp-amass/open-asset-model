// Copyright © by Jeff Foley 2017-2026. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestStreetAddressKey(t *testing.T) {
	want := "123 Some Road New York NY"
	loc := StreetAddress{Address: want}

	if got := loc.Key(); got != want {
		t.Errorf("StreetAddress.Key() = %v, want %v", got, want)
	}
}

func TestStreetAddressAssetType(t *testing.T) {
	var _ model.Asset = StreetAddress{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*StreetAddress)(nil) // Verify the pointer properly implements the  Asset interface.

	loc := StreetAddress{}
	if loc.AssetType() != model.StreetAddress {
		t.Errorf("Expected asset type %s but got %s", model.StreetAddress, loc.AssetType())
	}
}

func TestStreetAddressJSON(t *testing.T) {
	loc := StreetAddress{
		Address:        "123 Main St",
		Building:       "Building A",
		BuildingNumber: "123",
		StreetName:     "Main St",
		Unit:           "Apt 1",
		POBox:          "P.O. Box 145",
		City:           "Anytown",
		Locality:       "Anytown",
		Province:       "Anyregion",
		Country:        "US",
		PostalCode:     "12345",
		GLN:            1234567890123,
	}

	expectedJSON := `{"address":"123 Main St","building":"Building A","building_number":"123","street_name":"Main St","unit":"Apt 1","po_box":"P.O. Box 145","city":"Anytown","locality":"Anytown","province":"Anyregion","country":"US","postal_code":"12345","gln":1234567890123}`

	jsonData, err := loc.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s but got %s", expectedJSON, string(jsonData))
	}
}
