package network_test

import (
	"encoding/json"
	"testing"

	. "github.com/owasp-amass/open-asset-model/network"

	"github.com/stretchr/testify/assert"
)

func TestRIROrganization(t *testing.T) {
	t.Run("Test successful creation of RIROrganization with valid name, RIRId, and RIR", func(t *testing.T) {
		rirOrg := RIROrganization{Name: "ExampleOrg", RIRId: "e12345", RIR: "ARIN"}

		assert.Equal(t, "ExampleOrg", rirOrg.Name)
		assert.Equal(t, "e12345", rirOrg.RIRId)
		assert.Equal(t, "ARIN", rirOrg.RIR)
	})

	t.Run("Test successful JSON serialization of RIROrganization with valid name, RIRId, and RIR", func(t *testing.T) {
		rirOrg := RIROrganization{Name: "ExampleOrg", RIRId: "e12345", RIR: "ARIN"}

		jsonData, err := json.Marshal(rirOrg)

		assert.NoError(t, err)
		assert.Contains(t, string(jsonData), `"name":"ExampleOrg"`)
		assert.Contains(t, string(jsonData), `"rir_id":"e12345"`)
		assert.Contains(t, string(jsonData), `"rir":"ARIN"`)
	})
}
