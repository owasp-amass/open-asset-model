// Copyright © by Jeff Foley 2017-2026. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package open_asset_model

import "testing"

// A File records what it mentions (contains), who made it (author), the
// software that produced it (created_with) and its content hash
// (fingerprint). Where it was retrieved is the URL "file" relation only.
func TestFileRelations(t *testing.T) {
	for _, dst := range []AssetType{Account, ContactRecord, FQDN, File, GeoLocation, Identifier, IPAddress, Location,
		Netblock, Organization, Person, Phone, URL} {
		if !ValidRelationship(File, "contains", SimpleRelation, dst) {
			t.Errorf("File --contains--> %s is not valid", dst)
		}
	}
	for _, dst := range []AssetType{Organization, Person} {
		if !ValidRelationship(File, "author", SimpleRelation, dst) {
			t.Errorf("File --author--> %s is not valid", dst)
		}
	}
	for _, dst := range []AssetType{Product, ProductRelease} {
		if !ValidRelationship(File, "created_with", SimpleRelation, dst) {
			t.Errorf("File --created_with--> %s is not valid", dst)
		}
	}
	if !ValidRelationship(File, "fingerprint", SimpleRelation, Identifier) {
		t.Error("File --fingerprint--> Identifier is not valid")
	}
	if !ValidRelationship(URL, "file", SimpleRelation, File) {
		t.Error("URL --file--> File is not valid")
	}

	if ValidRelationship(File, "url", SimpleRelation, URL) {
		t.Error("File --url--> URL repeats URL --file--> File and was removed")
	}
	if ValidRelationship(File, "author", SimpleRelation, FQDN) {
		t.Error("a name cannot author a document")
	}
}

// Positions are reached from what is at them: a street address, an IP
// address's geolocation, a file that records one. A GeoLocation points
// nowhere.
func TestGeoLocationRelations(t *testing.T) {
	for _, src := range []AssetType{Location, IPAddress} {
		if !ValidRelationship(src, "geolocation", SimpleRelation, GeoLocation) {
			t.Errorf("%s --geolocation--> GeoLocation is not valid", src)
		}
	}
	if !ValidRelationship(File, "contains", SimpleRelation, GeoLocation) {
		t.Error("File --contains--> GeoLocation is not valid")
	}
	if rels := GetAssetOutgoingRelations(GeoLocation); len(rels) != 0 {
		t.Errorf("GeoLocation has outgoing relations %v", rels)
	}
}
