// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package domain

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestNetworkEndpointKey(t *testing.T) {
	want := "example.com:80"
	ne := NetworkEndpoint{Address: want}

	if got := ne.Key(); got != want {
		t.Errorf("NetworkEndpoint.Key() = %v, want %v", got, want)
	}
}

func TestNetworkEndpointAssetType(t *testing.T) {
	var _ model.Asset = NetworkEndpoint{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*NetworkEndpoint)(nil) // Verify the pointer properly implements the  Asset interface.

	sa := NetworkEndpoint{}
	want := model.NetworkEndpoint

	if got := sa.AssetType(); got != want {
		t.Errorf("NetworkEndpoint.AssetType() = %v, want %v", got, want)
	}
}

func TestNetworkEndpointJSON(t *testing.T) {
	sa := NetworkEndpoint{
		Address:  "example.com:80",
		Name:     "example.com",
		Port:     80,
		Protocol: "tcp",
	}

	// Test AssetType method
	if sa.AssetType() != model.NetworkEndpoint {
		t.Errorf("Expected asset type %s, but got %s", model.NetworkEndpoint, sa.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"address":"example.com:80","name":"example.com","port":80,"protocol":"tcp"}`
	json, err := sa.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
