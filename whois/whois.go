package whois

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// WHOIS represents the WHOIS information of a domain.
type WHOIS struct {
	// Type represents the type of the WHOIS information.
	Type string `json:"type"`

	// DomainStatus represents the status of the domain.
	DomainStatus []string `json:"domain_status,omitempty"`

	// WHOISServer represents the WHOIS server of the domain.
	WHOISServer string `json:"whois_server,omitempty"`

	// RegistryDomainID represents the registry domain ID of the domain.
	RegistryDomainID string `json:"registry_domain_id,omitempty"`

	// IANAid represents the IANA ID of the domain.
	IANAid string `json:"iana_id,omitempty"`

	// Description represents the description of the domain or org associated with a netblock.
	Description string `json:"description,omitempty"`

	// CreatedDate represents the creation date of the domain or org associated with a netblock.
	CreatedDate string `json:"created_at,omitempty"`

	// LastModifiedDate represents the last modified date of the domain.
	LastModifiedDate string `json:"last_modified,omitempty"`

	// ExpirationDate represents the expiration date of the domain.
	ExpirationDate string `json:"expiration_date,omitempty"`

	// DNSSEC represents the DNSSEC status of the domain.
	DNSSEC string `json:"dnssec,omitempty"`
}

// AssetType returns the asset type.
func (w WHOIS) AssetType() model.AssetType {
	return model.WHOIS
}

// JSON returns the JSON representation of the asset.
func (w WHOIS) JSON() ([]byte, error) {
	return json.Marshal(w)
}
