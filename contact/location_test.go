package contact_test

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/contact"
)

func TestLocation_AssetType(t *testing.T) {
	loc := Location{}
	if loc.AssetType() != model.Location {
		t.Errorf("Expected asset type %s but got %s", model.Location, loc.AssetType())
	}
}

func TestLocation_JSON(t *testing.T) {
	loc := Location{
		FormattedAddress: "123 Main St",
		BuildingNumber:   "123",
		StreetName:       "Main St",
		Unit:             "Apt 1",
		Building:         "Building A",
		Town:             "Anytown",
		Locality:         "Anytown",
		Region:           "Anyregion",
		CountryCode:      "US",
		PostalCode:       "12345",
	}

	expectedJSON := `{"formatted_address":"123 Main St","building_number":"123","street_name":"Main St","unit":"Apt 1","building":"Building A","town":"Anytown","locality":"Anytown","region":"Anyregion","country_code":"US","postal_code":"12345"}`

	jsonData, err := loc.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s but got %s", expectedJSON, string(jsonData))
	}
}
