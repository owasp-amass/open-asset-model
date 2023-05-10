package domain_test

import (
	"encoding/json"
	"testing"

	. "github.com/owasp-amass/open-asset-model/domain"

	"github.com/stretchr/testify/require"
)

func TestFQDN(t *testing.T) {
	t.Run("Test successful creation of FQDN with valid name and TLD", func(t *testing.T) {
		fqdn := FQDN{Name: "foo.example.com"}

		require.Equal(t, "foo.example.com", fqdn.Name)
	})

	t.Run("Test successful JSON serialization of FQDN with valid name and TLD", func(t *testing.T) {
		fqdn := FQDN{Name: "foo.example.com"}

		jsonData, err := json.Marshal(fqdn)

		require.NoError(t, err)
		require.JSONEq(t, `{"name":"foo.example.com"}`, string(jsonData))
	})

}
