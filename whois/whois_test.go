// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package whois

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestWHOIS_AssetType(t *testing.T) {
	w := WHOIS{}
	want := model.WHOIS

	if got := w.AssetType(); got != want {
		t.Errorf("WHOIS.AssetType() = %v, want %v", got, want)
	}
}
func TestWHOIS(t *testing.T) {
	whois := WHOIS{
		Type:                       "domain",
		Registrar:                  "Registrar",
		Domain:                     "example.com",
		Reseller:                   "Reseller",
		NameServers:                []string{"NameServer1", "NameServer2"},
		CreatedDate:                "2020-01-01",
		UpdatedDate:                "2021-01-01",
		ExpirationDate:             "2022-01-01",
		DomainStatus:               []string{"active", "clientTransferProhibited"},
		RegistryRegistrantID:       "12345",
		RegistryRegistrantName:     "RegistryRegistrantName",
		RegistryRegistrantOrg:      "RegistryRegistrantOrg",
		RegistryRegistrantLocation: "RegistryRegistrantLocation",
		RegistryRegistrantPhone:    "RegistryRegistrantPhone",
		RegistryRegistrantFax:      "RegistryRegistrantFax",
		RegistryRegistrantEmail:    "RegistryRegistrantEmail",
		RegistryDomainID:           "67890",
		RegistryBillingID:          "13579",
		RegistryBillingName:        "RegistryBillingName",
		RegistryBillingOrg:         "RegistryBillingOrg",
		RegistryBillingLocation:    "RegistryBillingLocation",
		RegistryBillingEmail:       "RegistryBillingEmail",
		RegistryAdminID:            "24680",
		RegistryAdminName:          "RegistryAdminName",
		RegistryAdminOrg:           "RegistryAdminOrg",
		RegistryAdminLocation:      "RegistryAdminLocation",
		RegistryAdminPhone:         "RegistryAdminPhone",
		RegistryAdminFax:           "RegistryAdminFax",
		RegistryAdminEmail:         "RegistryAdminEmail",
		RegistryTechID:             "13579",
		RegistryTechName:           "RegistryTechName",
		RegistryTechOrg:            "RegistryTechOrg",
		RegistryTechLocation:       "RegistryTechLocation",
		RegistryTechPhone:          "RegistryTechPhone",
		RegistryTechFax:            "RegistryTechFax",
		RegistryTechEmail:          "RegistryTechEmail",
		Description:                "Example domain",
		DNSSEC:                     "unsigned",
	}

	// Test AssetType method
	if whois.AssetType() != model.WHOIS {
		t.Errorf("Expected asset type %s, but got %s", model.WHOIS, whois.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"type":"domain","registrar":"Registrar","domain":"example.com","reseller":"Reseller","name_servers":["NameServer1","NameServer2"],"created_date":"2020-01-01","updated_date":"2021-01-01","expiration_date":"2022-01-01","domain_status":["active","clientTransferProhibited"],"registry_registrant_id":"12345","registry_registrant_name":"RegistryRegistrantName","registry_registrant_org":"RegistryRegistrantOrg","registry_registrant_location":"RegistryRegistrantLocation","registry_registrant_phone":"RegistryRegistrantPhone","registry_registrant_fax":"RegistryRegistrantFax","registry_registrant_email":"RegistryRegistrantEmail","registry_domain_id":"67890","registry_billing_id":"13579","registry_billing_name":"RegistryBillingName","registry_billing_org":"RegistryBillingOrg","registry_billing_location":"RegistryBillingLocation","registry_billing_email":"RegistryBillingEmail","registry_admin_id":"24680","registry_admin_name":"RegistryAdminName","registry_admin_org":"RegistryAdminOrg","registry_admin_location":"RegistryAdminLocation","registry_admin_phone":"RegistryAdminPhone","registry_admin_fax":"RegistryAdminFax","registry_admin_email":"RegistryAdminEmail","registry_tech_id":"13579","registry_tech_name":"RegistryTechName","registry_tech_org":"RegistryTechOrg","registry_tech_location":"RegistryTechLocation","registry_tech_phone":"RegistryTechPhone","registry_tech_fax":"RegistryTechFax","registry_tech_email":"RegistryTechEmail","description":"Example domain","dnssec":"unsigned"}`
	json, err := whois.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
