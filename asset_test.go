// Copyright © by Jeff Foley 2017-2026. All rights reserved.
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
		Account,
		AutnumRecord,
		AutonomousSystem,
		ContactRecord,
		DomainRecord,
		File,
		FQDN,
		FundsTransfer,
		GeoLocation,
		Identifier,
		IPAddress,
		IPNetRecord,
		Netblock,
		Organization,
		Person,
		Phone,
		Product,
		ProductRelease,
		Service,
		StorageBucket,
		StreetAddress,
		TLSCertificate,
		URL,
	}

	expectedTypes := []string{
		"Account",
		"AutnumRecord",
		"AutonomousSystem",
		"ContactRecord",
		"DomainRecord",
		"File",
		"FQDN",
		"FundsTransfer",
		"GeoLocation",
		"Identifier",
		"IPAddress",
		"IPNetRecord",
		"Netblock",
		"Organization",
		"Person",
		"Phone",
		"Product",
		"ProductRelease",
		"Service",
		"StorageBucket",
		"StreetAddress",
		"TLSCertificate",
		"URL",
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

func TestSecurityTxtServiceRelations(t *testing.T) {
	got := map[string]bool{}
	for _, r := range GetAssetOutgoingRelations(Service) {
		got[r] = true
	}
	for _, label := range []string{"security_policy", "security_canonical", "security_encryption",
		"security_acknowledgments", "security_hiring"} {
		if !got[label] {
			t.Errorf("Service is missing the %q relation", label)
		}
		if !ValidRelationship(Service, label, SimpleRelation, URL) {
			t.Errorf("%q should allow Service --SimpleRelation--> URL", label)
		}
	}
	if !got["security_contact"] || !ValidRelationship(Service, "security_contact", SimpleRelation, ContactRecord) {
		t.Error("security_contact should allow Service --SimpleRelation--> ContactRecord")
	}
}

func TestOnionLocationServiceRelation(t *testing.T) {
	found := false
	for _, r := range GetAssetOutgoingRelations(Service) {
		found = found || r == "onion_location"
	}
	if !found || !ValidRelationship(Service, "onion_location", SimpleRelation, URL) {
		t.Error("onion_location should allow Service --SimpleRelation--> URL")
	}
}

func TestResponseHeaderServiceRelations(t *testing.T) {
	for label, dest := range map[string]AssetType{"content_source": FQDN, "allowed_origin": FQDN,
		"report_endpoint": URL} {
		if !ValidRelationship(Service, label, SimpleRelation, dest) {
			t.Errorf("%q should allow Service --SimpleRelation--> %s", label, dest)
		}
	}
}
