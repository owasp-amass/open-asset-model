// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package registration

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestDomainRecordKey(t *testing.T) {
	want := "owasp.org"
	dr := DomainRecord{Domain: want}

	if got := dr.Key(); got != want {
		t.Errorf("DomainRecord.Key() = %v, want %v", got, want)
	}
}

func TestDomainRecordAssetType(t *testing.T) {
	var _ model.Asset = DomainRecord{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*DomainRecord)(nil) // Verify the pointer properly implements the  Asset interface.

	w := DomainRecord{}
	want := model.DomainRecord

	if got := w.AssetType(); got != want {
		t.Errorf("DomainRecord.AssetType() = %v, want %v", got, want)
	}
}
func TestDomainRecord(t *testing.T) {
	record := DomainRecord{
		Domain:         "example.com",
		CreatedDate:    "2020-01-01",
		UpdatedDate:    "2021-01-01",
		ExpirationDate: "2022-01-01",
		Status:         []string{"active", "clientTransferProhibited"},
		DNSSEC:         true,
	}

	// Test AssetType method
	if record.AssetType() != model.DomainRecord {
		t.Errorf("Expected asset type %s, but got %s", model.DomainRecord, record.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"domain":"example.com","created_date":"2020-01-01","updated_date":"2021-01-01","expiration_date":"2022-01-01","status":["active","clientTransferProhibited"],"dnssec":true}`
	json, err := record.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
