package contact

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// This struct represents the phone number, whether it is fax, mobile, or home number linked to the possible asset
type Phone struct {
	// Type is the type of phone number.
	Type string `json:"type"`
	// Raw is the raw phone number.
	Raw string `json:"raw"`
	// E164 is the phone number in E.164 format.
	E164 string `json:"e164"`
	// CountryAbbrev is the country abbreviation of the phone number.
	CountryAbbrev string `json:"country_abbrev"`
	// CountryCode is the country code of the phone number.
	CountryCode int `json:"country_code"`
	// SubscriberNumber is the subscriber number of the phone number.
	SubscriberNumber string `json:"subscriber_number"`
	// Ext is the extension of the phone number.
	Ext string `json:"ext,omitempty"`
}

// AssetType returns the asset type.
func (p Phone) AssetType() model.AssetType {
	return model.Phone
}

// JSON returns the JSON representation of the asset.
func (p Phone) JSON() ([]byte, error) {
	return json.Marshal(p)
}
