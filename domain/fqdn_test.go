// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package domain

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestFQDNKey(t *testing.T) {
	want := "example.com"
	fqdn := FQDN{Name: want}

	if got := fqdn.Key(); got != want {
		t.Errorf("FQDN.Key() = %v, want %v", got, want)
	}
}

func TestFQDNImplementsAsset(t *testing.T) {
	var _ model.Asset = FQDN{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*FQDN)(nil) // Verify *FQDN properly implements the  Asset interface.
}

func TestFQDN(t *testing.T) {
	t.Run("Test successful creation of FQDN with valid name and TLD", func(t *testing.T) {
		fqdn := FQDN{Name: "foo.example.com"}

		require.Equal(t, "foo.example.com", fqdn.Name)
		require.Equal(t, fqdn.AssetType(), model.FQDN)
	})

	t.Run("Test successful JSON serialization of FQDN with valid name and TLD", func(t *testing.T) {
		fqdn := FQDN{Name: "foo.example.com"}

		jsonData, err := fqdn.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"name":"foo.example.com"}`, string(jsonData))
	})
}
