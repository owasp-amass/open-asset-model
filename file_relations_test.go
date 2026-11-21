// Copyright © by Jeff Foley 2017-2026. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package open_asset_model

import "testing"

// A File records what it mentions (contains), who made it (author), the
// software that produced it (created_with) and its content hash
// (fingerprint). Where it was retrieved is the URL "file" relation only.
func TestFileRelations(t *testing.T) {
	for _, dst := range []AssetType{Account, ContactRecord, FQDN, File, Identifier, IPAddress, Location, Netblock,
		Organization, Person, Phone, URL} {
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
