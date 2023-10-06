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
		Type:            "domain",
		CreatedDate:     "2022-01-01",
		UpdatedDate:     "2022-01-02",
		ExpirationDate:  "2023-01-01",
		DomainStatus:    []string{"clientTransferProhibited", "serverTransferProhibited"},
		RegistryAdminID: "ADMIN123",
		Description:     "Example Domain",
		DNSSEC:          "unsigned",
	}

	want, err := json.Marshal(w)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	if got, err := w.JSON(); !reflect.DeepEqual(got, want) || err != nil {
		t.Errorf("WHOIS.JSON() = %v, %v; want %v, nil", got, err, want)
	}
}
