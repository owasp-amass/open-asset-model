package people

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Person represents a person's information.
type Person struct {
	// FullName is the full name of the person.
	FullName string `json:"full_name"`
	// FirstName is the first name of the person.
	FirstName string `json:"first_name"`
	// MiddleName is the middle name of the person.
	MiddleName string `json:"middle_name"`
	// FamilyName is the family name of the person.
	FamilyName string `json:"family_name"`
	// BirthCountry is the country where the person was born.
	BirthCountry string `json:"birth_country"`
	// DateOfBirth is the date of birth of the person.
	DateOfBirth string `json:"date_of_birth"`
}

// AssetType returns the asset type.
func (n Person) AssetType() model.AssetType {
	return model.Person
}

// JSON returns the JSON representation of the asset.
func (n Person) JSON() ([]byte, error) {
	return json.Marshal(n)
}
