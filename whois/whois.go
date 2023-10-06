package whois

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// WHOIS represents the WHOIS information of a domain.
type WHOIS struct {
	// Type represents the type of the WHOIS information.
	Type string `json:"type"`

	// CreatedDate represents the creation date of the domain or org associated with a netblock.
	CreatedDate string `json:"created_date,omitempty"`

	// LastModifiedDate represents the last modified date of the domain.
	UpdatedDate string `json:"updated_date,omitempty"`

	// ExpirationDate represents the expiration date of the domain.
	ExpirationDate string `json:"expiration_date,omitempty"`

	// DomainStatus represents the status of the domain.
	DomainStatus []string `json:"domain_status,omitempty"`

	// RegistryRegistrantID represents the registry registrant ID of the domain.
	RegistryRegistrantID string `json:"registry_registrant_id,omitempty"`

	// RegistryDomainID represents the registry domain ID of the domain.
	RegistryDomainID string `json:"registry_domain_id,omitempty"`

	// RegistryAdminID represents the registry admin ID of the domain.
	RegistryAdminID string `json:"registry_admin_id,omitempty"`

	// RegistryTechID represents the registry tech ID of the domain.
	RegistryTechID string `json:"registry_tech_id,omitempty"`

	// Description represents the description of the domain or org associated with a netblock.
	Description string `json:"description,omitempty"`

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
