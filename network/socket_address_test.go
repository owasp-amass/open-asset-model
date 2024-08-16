// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"net/netip"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestSocketAddressAssetType(t *testing.T) {
	sa := SocketAddress{}
	want := model.SocketAddress

	if got := sa.AssetType(); got != want {
		t.Errorf("Port.AssetType() = %v, want %v", got, want)
	}
}

func TestSocketAddressJSON(t *testing.T) {
	sa := SocketAddress{
		Address:   netip.MustParseAddrPort("192.168.1.1:80"),
		IPAddress: netip.MustParseAddr("192.168.1.1"),
		Port:      80,
		Protocol:  "tcp",
	}

	// Test AssetType method
	if sa.AssetType() != model.SocketAddress {
		t.Errorf("Expected asset type %s, but got %s", model.SocketAddress, sa.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"address":"192.168.1.1:80","ip_address":"192.168.1.1","port":80,"protocol":"tcp"}`
	json, err := sa.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
