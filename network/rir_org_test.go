// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestRIROrganizationImplementsAsset(t *testing.T) {
	var _ model.Asset = RIROrganization{}       // Verify that RIROrganization implements Asset interface
	var _ model.Asset = (*RIROrganization)(nil) // Verify that *RIROrganization implements Asset interface.
}

func TestRIROrganization(t *testing.T) {
	t.Run("Test successful creation of RIROrganization with valid name, RIRId, and RIR", func(t *testing.T) {
		rirOrg := RIROrganization{Name: "ExampleOrg", RIRId: "e12345", RIR: "ARIN"}

		require.Equal(t, "ExampleOrg", rirOrg.Name)
		require.Equal(t, "e12345", rirOrg.RIRId)
		require.Equal(t, "ARIN", rirOrg.RIR)
		require.Equal(t, rirOrg.AssetType(), model.RIROrg)
	})

	t.Run("Test successful JSON serialization of RIROrganization with valid name, RIRId, and RIR", func(t *testing.T) {
		rirOrg := RIROrganization{Name: "ExampleOrg", RIRId: "e12345", RIR: "ARIN"}

		jsonData, err := rirOrg.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"name":"ExampleOrg","rir_id":"e12345","rir":"ARIN"}`, string(jsonData))
	})
}
