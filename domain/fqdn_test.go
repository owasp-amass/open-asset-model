package domain_test

import (
	"testing"

	openAssetModel "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/domain"

	"github.com/stretchr/testify/require"
)

func TestFQDNImplementsAsset(t *testing.T) {
	var _ openAssetModel.Asset = FQDN{}       // Verify that FQDN implements Asset interface
	var _ openAssetModel.Asset = (*FQDN)(nil) // Verify that *FQDN implements Asset interface.
}

func TestFQDN(t *testing.T) {
	t.Run("Test successful creation of FQDN with valid name and TLD", func(t *testing.T) {
		fqdn := FQDN{Name: "foo.example.com"}

		require.Equal(t, "foo.example.com", fqdn.Name)
		require.Equal(t, fqdn.AssetType(), openAssetModel.FQDN)
	})

	t.Run("Test successful JSON serialization of FQDN with valid name and TLD", func(t *testing.T) {
		fqdn := FQDN{Name: "foo.example.com"}

		jsonData, err := fqdn.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"name":"foo.example.com"}`, string(jsonData))
	})
}
