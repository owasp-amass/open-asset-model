// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package registration

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestAutnumRecordAssetType(t *testing.T) {
	w := AutnumRecord{}
	want := model.AutnumRecord

	if got := w.AssetType(); got != want {
		t.Errorf("AutnumRecord.AssetType() = %v, want %v", got, want)
	}
}
func TestAutnumRecord(t *testing.T) {
	record := AutnumRecord{
		Number:      26808,
		Handle:      "AS26808",
		Name:        "UTICA-COLLEGE",
		WhoisServer: "whois.arin.net",
		CreatedDate: "2002-11-25 22:25:46",
		UpdatedDate: "2021-05-03 17:59:17",
		Status:      []string{"active"},
	}

	// Test AssetType method
	if record.AssetType() != model.AutnumRecord {
		t.Errorf("Expected asset type %s, but got %s", model.AutnumRecord, record.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"number":26808,"handle":"AS26808","name":"UTICA-COLLEGE","whois_server":"whois.arin.net","created_date":"2002-11-25 22:25:46","updated_date":"2021-05-03 17:59:17","status":["active"]}`
	json, err := record.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
