// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package open_asset_model

import (
	"testing"
)

type MockAsset struct {
	assetType AssetType
}

func (a MockAsset) AssetType() AssetType {
	return a.assetType
}

func TestAssetInterface(t *testing.T) {
	mockAsset := MockAsset{assetType: IPAddress}

	if mockAsset.AssetType() != IPAddress {
		t.Error("AssetType() does not return the expected value")
	}
}

func TestAssetTypeConstants(t *testing.T) {
	assetTypes := []AssetType{
		IPAddress,
		Netblock,
		AutonomousSystem,
		FQDN,
		DomainRecord,
		AutnumRecord,
		Location,
		Phone,
		EmailAddress,
		Person,
		Organization,
		URL,
		TLSCertificate,
		ContactRecord,
	}

	expectedTypes := []string{
		"IPAddress",
		"Netblock",
		"AutonomousSystem",
		"FQDN",
		"DomainRecord",
		"AutnumRecord",
		"Location",
		"Phone",
		"EmailAddress",
		"Person",
		"Organization",
		"URL",
		"TLSCertificate",
		"ContactRecord",
	}

	if len(assetTypes) != len(expectedTypes) {
		t.Error("Number of asset types does not match the expected number")
	}

	for i, assetType := range assetTypes {
		if string(assetType) != expectedTypes[i] {
			t.Errorf("Asset type %s does not match the expected value %s", assetType, expectedTypes[i])
		}
	}
}
