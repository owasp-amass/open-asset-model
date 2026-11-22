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

func TestBugBountyProgramRelations(t *testing.T) {
	for _, dest := range []AssetType{FQDN, Netblock, URL} {
		if !ValidRelationship(Account, "in_scope", SimpleRelation, dest) {
			t.Errorf("in_scope should allow Account --SimpleRelation--> %s", dest)
		}
	}
	if !ValidRelationship(Account, "program_page", SimpleRelation, URL) {
		t.Error("program_page should allow Account --SimpleRelation--> URL")
	}
}

func TestURLContactRelation(t *testing.T) {
	if !ValidRelationship(URL, "contact", SimpleRelation, ContactRecord) {
		t.Error("contact should allow URL --SimpleRelation--> ContactRecord")
	}
}

func TestAccountContactRelation(t *testing.T) {
	if !ValidRelationship(Account, "contact", SimpleRelation, ContactRecord) {
		t.Error("contact should allow Account --SimpleRelation--> ContactRecord")
	}
}

func TestPackageNamespaceRelations(t *testing.T) {
	if !ValidRelationship(Account, "package", SimpleRelation, Product) {
		t.Error("package should allow Account --SimpleRelation--> Product")
	}
	if !ValidRelationship(Account, "namespace_domain", SimpleRelation, FQDN) {
		t.Error("namespace_domain should allow Account --SimpleRelation--> FQDN")
	}
}

func TestRoutingRegistryNetblockRelations(t *testing.T) {
	if !ValidRelationship(Netblock, "irr_maintainer", SimpleRelation, Identifier) {
		t.Error("irr_maintainer should allow Netblock --SimpleRelation--> Identifier")
	}
	if !ValidRelationship(Netblock, "irr_origin", SimpleRelation, AutonomousSystem) {
		t.Error("irr_origin should allow Netblock --SimpleRelation--> AutonomousSystem")
	}
	if !ValidRelationship(Netblock, "rpki_origin", SimpleRelation, AutonomousSystem) {
		t.Error("rpki_origin should allow Netblock --SimpleRelation--> AutonomousSystem")
	}
}

func TestSOAAndSPFRangeRelations(t *testing.T) {
	if !ValidRelationship(FQDN, "soa_rname", SimpleRelation, Identifier) {
		t.Error("soa_rname should allow FQDN --SimpleRelation--> Identifier")
	}
	for _, dest := range []AssetType{FQDN, Netblock, IPAddress} {
		if !ValidRelationship(FQDN, "spf_authorizes", SimpleRelation, dest) {
			t.Errorf("spf_authorizes should allow FQDN --SimpleRelation--> %s", dest)
		}
	}
}
