package network_test

import (
	"encoding/json"
	"testing"

	. "github.com/owasp-amass/open-asset-model/network"

	"github.com/stretchr/testify/require"
)

func TestRIROrganization(t *testing.T) {
	t.Run("Test successful creation of RIROrganization with valid name, RIRId, and RIR", func(t *testing.T) {
		rirOrg := RIROrganization{Name: "ExampleOrg", RIRId: "e12345", RIR: "ARIN"}

		require.Equal(t, "ExampleOrg", rirOrg.Name)
		require.Equal(t, "e12345", rirOrg.RIRId)
		require.Equal(t, "ARIN", rirOrg.RIR)
	})

	t.Run("Test successful JSON serialization of RIROrganization with valid name, RIRId, and RIR", func(t *testing.T) {
		rirOrg := RIROrganization{Name: "ExampleOrg", RIRId: "e12345", RIR: "ARIN"}

		jsonData, err := json.Marshal(rirOrg)

		require.NoError(t, err)
		require.JSONEq(t, `{"name":"ExampleOrg","rir_id":"e12345","rir":"ARIN"}`, string(jsonData))
	})
}
