// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"strconv"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestAutonomousSystemKey(t *testing.T) {
	want := 26808
	a := AutonomousSystem{Number: want}

	if got := a.Key(); got != strconv.Itoa(want) {
		t.Errorf("AutonomousSystem.Key() = %v, want %v", got, want)
	}
}

func TestAutonomousSystemAsset(t *testing.T) {
	var _ model.Asset = AutonomousSystem{}       // Verify that AutonomousSystem implements Asset interface
	var _ model.Asset = (*AutonomousSystem)(nil) // Verify that *AutonomousSystem implements Asset interface.
}

func TestAutonomousSystem(t *testing.T) {
	t.Run("Test successful creation of AutonomousSystem with valid AS number", func(t *testing.T) {
		as := AutonomousSystem{Number: 64496}

		require.Equal(t, 64496, as.Number)
		require.Equal(t, as.AssetType(), model.AutonomousSystem)
	})

	t.Run("Test successful JSON serialization of AutonomousSystem with valid AS number", func(t *testing.T) {
		as := AutonomousSystem{Number: 64496}

		jsonData, err := as.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"number":64496}`, string(jsonData))
	})
}
