// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package fingerprint

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestFingerprintKey(t *testing.T) {
	want := "12345"
	fp := Fingerprint{Value: want}

	if got := fp.Key(); got != want {
		t.Errorf("Fingerprint.Key() = %v, want %v", got, want)
	}
}

func TestFingerprintAssetType(t *testing.T) {
	var _ model.Asset = Fingerprint{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*Fingerprint)(nil) // Verify the pointer properly implements the  Asset interface.

	fp := Fingerprint{
		Value: "example",
		Type:  "example",
	}
	want := model.Fingerprint

	if got := fp.AssetType(); got != want {
		t.Errorf("Fingerprint.AssetType() = %v, want %v", got, want)
	}
}

func TestFingerprintJSON(t *testing.T) {
	fp := Fingerprint{
		Value: "example",
		Type:  "example",
	}

	// Test AssetType method
	if fp.AssetType() != model.Fingerprint {
		t.Errorf("Expected asset type %s, but got %s", model.Fingerprint, fp.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"value":"example","type":"example"}`
	jsonData, err := fp.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(jsonData))
	}
}
