package whois_test

import (
	"encoding/json"
	"reflect"
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

func TestWHOIS_JSON(t *testing.T) {
	w := WHOIS{
		Type:                  "domain",
		Registrar_WHOISServer: "whois.example.com",
		Registrar_URL:         "http://www.example.com",
		CreatedDate:           "2022-01-01",
		LastModifiedDate:      "2022-01-02",
		ExpirationDate:        "2023-01-01",
		Reseller:              "Example Reseller",
		DomainStatus:          []string{"clientTransferProhibited", "serverTransferProhibited"},
		RegistryRegistrantID:  "REGISTRANT123",
		RegistryDomainID:      "DOMAIN123",
		RegistryAdminID:       "ADMIN123",
		RegistryTechID:        "TECH123",
		Description:           "Example Domain",
		DNSSEC:                "unsigned",
	}

	want, err := json.Marshal(w)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	if got, err := w.JSON(); !reflect.DeepEqual(got, want) || err != nil {
		t.Errorf("WHOIS.JSON() = %v, %v; want %v, nil", got, err, want)
	}
}
