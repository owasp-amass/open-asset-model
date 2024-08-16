// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package whois

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// WHOIS represents the WHOIS information of a domain.
type WHOIS struct {
	// Type represents the type of the WHOIS information.
	Type string `json:"type"`

	// Registrar represents the registrar of a domain name.
	Registrar string `json:"registrar,omitempty"`

	// Domain represents the domain name.
	Domain string `json:"domain,omitempty"`

	// Reseller represents the reseller of the domain.
	Reseller string `json:"reseller,omitempty"`

	// NameServers represents the name servers of the domain.
	NameServers []string `json:"name_servers,omitempty"`

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

	// RegistryRegistrantName represents the registry registrant name of the domain.
	RegistryRegistrantName string `json:"registry_registrant_name,omitempty"`

	// RegistryRegistrantOrg represents the registry registrant org of the domain.
	RegistryRegistrantOrg string `json:"registry_registrant_org,omitempty"`

	// RegistryRegistrantLocation represents the registry registrant location of the domain.
	RegistryRegistrantLocation string `json:"registry_registrant_location,omitempty"`

	// RegistryRegistrantPhone represents the registry registrant phone of the domain.
	RegistryRegistrantPhone string `json:"registry_registrant_phone,omitempty"`

	// RegistryRegistrantFax represents the registry registrant fax of the domain.
	RegistryRegistrantFax string `json:"registry_registrant_fax,omitempty"`

	// RegistryRegistrantEmail represents the registry registrant email of the domain.
	RegistryRegistrantEmail string `json:"registry_registrant_email,omitempty"`

	// RegistryDomainID represents the registry domain ID of the domain.
	RegistryDomainID string `json:"registry_domain_id,omitempty"`

	// RegistryBillingID represents the registry billing ID of the domain.
	RegistryBillingID string `json:"registry_billing_id,omitempty"`

	// RegistryBillingName represents the registry billing name of the domain.
	RegistryBillingName string `json:"registry_billing_name,omitempty"`

	// RegistryBillingOrg represents the registry billing org of the domain.
	RegistryBillingOrg string `json:"registry_billing_org,omitempty"`

	// RegistryBillingLocation represents the registry billing location of the domain.
	RegistryBillingLocation string `json:"registry_billing_location,omitempty"`

	// RegistryBillingPhone represents the registry billing phone of the domain.
	RegistryBillingPhone string `json:"registry_billing_phone,omitempty"`

	// RegistryBillingFax represents the registry billing fax of the domain.
	RegistryBillingFax string `json:"registry_billing_fax,omitempty"`

	// RegistryBillingEmail represents the registry billing email of the domain.
	RegistryBillingEmail string `json:"registry_billing_email,omitempty"`

	// RegistryAdminID represents the registry admin ID of the domain.
	RegistryAdminID string `json:"registry_admin_id,omitempty"`

	// RegistryAdminName represents the registry admin name of the domain.
	RegistryAdminName string `json:"registry_admin_name,omitempty"`

	// RegistryAdminOrg represents the registry admin org of the domain.
	RegistryAdminOrg string `json:"registry_admin_org,omitempty"`

	// RegistryAdminLocation represents the registry admin location of the domain.
	RegistryAdminLocation string `json:"registry_admin_location,omitempty"`

	// RegistryAdminPhone represents the registry admin phone of the domain.
	RegistryAdminPhone string `json:"registry_admin_phone,omitempty"`

	// RegistryAdminFax represents the registry admin fax of the domain.
	RegistryAdminFax string `json:"registry_admin_fax,omitempty"`

	// RegistryAdminEmail represents the registry admin email of the domain.
	RegistryAdminEmail string `json:"registry_admin_email,omitempty"`

	// RegistryTechID represents the registry tech ID of the domain.
	RegistryTechID string `json:"registry_tech_id,omitempty"`

	// RegistryTechName represents the registry tech name of the domain.
	RegistryTechName string `json:"registry_tech_name,omitempty"`

	// RegistryTechOrg represents the registry tech org of the domain.
	RegistryTechOrg string `json:"registry_tech_org,omitempty"`

	// RegistryTechLocation represents the registry tech location of the domain.
	RegistryTechLocation string `json:"registry_tech_location,omitempty"`

	// RegistryTechPhone represents the registry tech phone of the domain.
	RegistryTechPhone string `json:"registry_tech_phone,omitempty"`

	// RegistryTechFax represents the registry tech fax of the domain.
	RegistryTechFax string `json:"registry_tech_fax,omitempty"`

	// RegistryTechEmail represents the registry tech email of the domain.
	RegistryTechEmail string `json:"registry_tech_email,omitempty"`

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
