package whois_test

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/whois"
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
		Type:                 "domain",
		CreatedDate:          "2020-01-01",
		UpdatedDate:          "2021-01-01",
		ExpirationDate:       "2022-01-01",
		DomainStatus:         []string{"active", "clientTransferProhibited"},
		RegistryRegistrantID: "12345",
		RegistryDomainID:     "67890",
		RegistryBillingID:    "13579",
		RegistryAdminID:      "24680",
		RegistryTechID:       "13579",
		Description:          "Example domain",
		DNSSEC:               "unsigned",
	}

	// Test AssetType method
	if whois.AssetType() != model.WHOIS {
		t.Errorf("Expected asset type %s, but got %s", model.WHOIS, whois.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"type":"domain","created_date":"2020-01-01","updated_date":"2021-01-01","expiration_date":"2022-01-01","domain_status":["active","clientTransferProhibited"],"registry_registrant_id":"12345","registry_domain_id":"67890","registry_billing_id":"13579","registry_admin_id":"24680","registry_tech_id":"13579","description":"Example domain","dnssec":"unsigned"}`
	json, err := whois.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
