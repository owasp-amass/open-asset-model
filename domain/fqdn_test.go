package domain_test

import (
	"encoding/json"
	"testing"

	. "github.com/owasp-amass/open-asset-model/domain"

	"github.com/stretchr/testify/assert"
)

func TestFQDN(t *testing.T) {
	t.Run("Test successful creation of FQDN with valid name and TLD", func(t *testing.T) {
		fqdn := FQDN{Name: "foo.example.com"}

		assert.Equal(t, "foo.example.com", fqdn.Name)
	})

	t.Run("Test successful JSON serialization of FQDN with valid name and TLD", func(t *testing.T) {
		fqdn := FQDN{Name: "foo.example.com"}

		jsonData, err := json.Marshal(fqdn)

		assert.NoError(t, err)
		assert.Contains(t, string(jsonData), `"name":"foo.example.com"`)
	})

}
