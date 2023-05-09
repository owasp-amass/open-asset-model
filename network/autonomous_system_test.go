package network_test

import (
	"encoding/json"
	"testing"

	. "github.com/owasp-amass/open-asset-model/network"

	"github.com/stretchr/testify/assert"
)

func TestAutonomousSystem(t *testing.T) {
	t.Run("Test successful creation of AutonomousSystem with valid AS number", func(t *testing.T) {
		as := AutonomousSystem{Number: 64496}

		assert.Equal(t, 64496, as.Number)
	})

	t.Run("Test successful JSON serialization of AutonomousSystem with valid AS number", func(t *testing.T) {
		as := AutonomousSystem{Number: 64496}

		jsonData, err := json.Marshal(as)

		assert.NoError(t, err)
		assert.Contains(t, string(jsonData), `"number":64496`)
	})
}
