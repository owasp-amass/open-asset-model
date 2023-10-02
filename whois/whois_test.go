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
		Type:             "domain",
		DomainStatus:     []string{"clientTransferProhibited", "clientUpdateProhibited", "clientDeleteProhibited"},
		WHOISServer:      "whois.markmonitor.com",
		RegistryDomainID: "123456789_DOMAIN_COM-VRSN",
		IANAid:           "292",
		Description:      "This is a test domain",
		CreatedDate:      "2021-01-01T00:00:00Z",
		LastModifiedDate: "2021-01-02T00:00:00Z",
		ExpirationDate:   "2022-01-01T00:00:00Z",
		DNSSEC:           "unsigned",
	}
	want, _ := json.Marshal(w)

	if got, err := w.JSON(); !reflect.DeepEqual(got, want) || err != nil {
		t.Errorf("WHOIS.JSON() = %v, want %v, error: %v", got, want, err)
	}
}
