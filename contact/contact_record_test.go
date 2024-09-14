// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestContactRecordKey(t *testing.T) {
	want := "https://owasp.org/contacts"
	cr := ContactRecord{DiscoveredAt: want}

	if got := cr.Key(); got != want {
		t.Errorf("ContactRecord.Key() = %v, want %v", got, want)
	}
}

func TestContactRecordAssetType(t *testing.T) {
	var _ model.Asset = ContactRecord{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*ContactRecord)(nil) // Verify the pointer properly implements the  Asset interface.

	cr := ContactRecord{}
	if cr.AssetType() != model.ContactRecord {
		t.Errorf("Expected asset type %s but got %s", model.ContactRecord, cr.AssetType())
	}
}

func TestContactRecordJSON(t *testing.T) {
	cr := ContactRecord{DiscoveredAt: "https://owasp.org"}

	expectedJSON := `{"discovered_at":"https://owasp.org"}`

	jsonData, err := cr.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s but got %s", expectedJSON, string(jsonData))
	}
}
