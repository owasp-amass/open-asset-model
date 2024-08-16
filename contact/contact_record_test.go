// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestContactRecordAssetType(t *testing.T) {
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
