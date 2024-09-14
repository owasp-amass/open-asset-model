// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package service

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestServiceKey(t *testing.T) {
	want := "12345"
	serv := Service{Identifier: want}

	if got := serv.Key(); got != want {
		t.Errorf("Service.Key() = %v, want %v", got, want)
	}
}

func TestServiceAssetType(t *testing.T) {
	var _ model.Asset = Service{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*Service)(nil) // Verify the pointer properly implements the  Asset interface.

	s := Service{}
	want := model.Service

	if got := s.AssetType(); got != want {
		t.Errorf("Service.AssetType() = %v, want %v", got, want)
	}
}

func TestServiceJSON(t *testing.T) {
	s := Service{
		Identifier: "12345",
		Version:    "v1.0.1",
	}

	// test AssetType method
	if s.AssetType() != model.Service {
		t.Errorf("Expected asset type %s, but got %s", model.Service, s.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"identifier":"12345","version":"v1.0.1"}`
	json, err := s.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
