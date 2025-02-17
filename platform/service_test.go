// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package platform

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestServiceKey(t *testing.T) {
	want := "222333444"
	serv := Service{
		ID:   want,
		Type: "HTTP",
	}

	if got := serv.Key(); got != want {
		t.Errorf("Service.Key() = %v, want %v", got, want)
	}
}

func TestServiceAssetType(t *testing.T) {
	var _ model.Asset = Service{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*Service)(nil) // Verify the pointer properly implements the Asset interface

	s := Service{}
	want := model.Service

	if got := s.AssetType(); got != want {
		t.Errorf("Service.AssetType() = %v, want %v", got, want)
	}
}

func TestServiceJSON(t *testing.T) {
	s := Service{
		ID:         "222333444",
		Type:       "HTTP",
		Output:     "Hello",
		OutputLen:  5,
		Attributes: map[string][]string{"server": {"nginx-1.26.0"}},
	}

	// test AssetType method
	if s.AssetType() != model.Service {
		t.Errorf("Expected asset type %s, but got %s", model.Service, s.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"unique_id":"222333444","service_type":"HTTP","output":"Hello","output_length":5,"attributes":{"server":["nginx-1.26.0"]}}`
	json, err := s.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
