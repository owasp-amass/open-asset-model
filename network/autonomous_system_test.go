package network_test

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/network"
	"github.com/stretchr/testify/require"
)

func TestAutonomousSystemImplementsAsset(t *testing.T) {
	var _ model.Asset = AutonomousSystem{}       // Verify that AutonomousSystem implements Asset interface
	var _ model.Asset = (*AutonomousSystem)(nil) // Verify that *AutonomousSystem implements Asset interface.
}

func TestAutonomousSystem(t *testing.T) {
	t.Run("Test successful creation of AutonomousSystem with valid AS number", func(t *testing.T) {
		as := AutonomousSystem{Number: 64496}

		require.Equal(t, 64496, as.Number)
		require.Equal(t, as.AssetType(), model.ASN)
	})

	t.Run("Test successful JSON serialization of AutonomousSystem with valid AS number", func(t *testing.T) {
		as := AutonomousSystem{Number: 64496}

		jsonData, err := as.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"number":64496}`, string(jsonData))
	})
}
