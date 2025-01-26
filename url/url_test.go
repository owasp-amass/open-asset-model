// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package url

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestURLKey(t *testing.T) {
	want := "https://owasp.org"
	u := URL{Raw: want}

	if got := u.Key(); got != want {
		t.Errorf("URL.Key() = %v, want %v", got, want)
	}
}

func TestURLAssetType(t *testing.T) {
	var _ model.Asset = URL{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*URL)(nil) // Verify the pointer properly implements the  Asset interface.

	u := URL{}
	want := model.URL

	if got := u.AssetType(); got != want {
		t.Errorf("URL.AssetType() = %v, want %v", got, want)
	}
}

func TestURLJSON(t *testing.T) {
	u := URL{
		Raw:      "http://user:pass@example.com:8080/path?option1=value1&option2=value2#fragment",
		Scheme:   "http",
		Username: "user",
		Password: "pass",
		Host:     "example.com",
		Port:     8080,
		Path:     "/path",
		Options:  "option1=value1&option2=value2",
		Fragment: "fragment",
	}

	// Test AssetType method
	if u.AssetType() != model.URL {
		t.Errorf("Expected asset type %s, but got %s", model.URL, u.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"url":"http://user:pass@example.com:8080/path?option1=value1&option2=value2#fragment","scheme":"http","username":"user","password":"pass","host":"example.com","port":8080,"path":"/path","options":"option1=value1&option2=value2","fragment":"fragment"}`
	json, err := u.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
