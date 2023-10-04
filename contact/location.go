package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// This struct represents the location linked to the possible asset
type Location struct {
	// FormattedAddress is the human-readable address of the location.
	FormattedAddress string `json:"formatted_address,omitempty"`

	// BuildingNumber is the number of the building at the location.
	BuildingNumber string `json:"building_number,omitempty"`

	// StreetName is the name of the street at the location.
	StreetName string `json:"street_name,omitempty"`

	// Unit is the unit number of the location.
	Unit string `json:"unit,omitempty"`

	// Building is the name or number of the building at the location.
	Building string `json:"building,omitempty"`

	// Town is the name of the town or city at the location.
	Town string `json:"town,omitempty"`

	// Locality is the name of the locality at the location.
	Locality string `json:"locality,omitempty"`

	// Region is the name of the region or state at the location.
	Region string `json:"region,omitempty"`

	// CountryCode is the ISO 3166-1 alpha-2 country code of the location.
	CountryCode string `json:"country_code,omitempty"`

	// PostalCode is the postal code of the location.
	PostalCode string `json:"postal_code,omitempty"`

	// LatLong is the latitude and longitude of the location.
	LatLong [2]float64 `json:"lat_long,omitempty"`
}

// AssetType returns the asset type.
func (a Location) AssetType() model.AssetType {
	return model.Location
}

// JSON returns the JSON representation of the asset.
func (a Location) JSON() ([]byte, error) {
	return json.Marshal(a)
}
