// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package domain

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestFQDNImplementsAsset(t *testing.T) {
	var _ model.Asset = FQDN{}       // Verify that FQDN implements Asset interface
	var _ model.Asset = (*FQDN)(nil) // Verify that *FQDN implements Asset interface.
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
