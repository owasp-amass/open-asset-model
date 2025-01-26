// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package id

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestIdentifierKey(t *testing.T) {
	want := "Legal Entity Identifier:549300XMYB546ZI1F126"
	i := Identifier{ID: "549300XMYB546ZI1F126", Type: "Legal Entity Identifier"}

	if got := i.Key(); got != want {
		t.Errorf("Identifier.Key() = %v, want %v", got, want)
	}
}

func TestIdentifierAssetType(t *testing.T) {
	var _ model.Asset = Identifier{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*Identifier)(nil) // Verify the pointer properly implements the  Asset interface.

	i := Identifier{}
	expected := model.Identifier
	actual := i.AssetType()

	if actual != expected {
		t.Errorf("Expected asset type %v but got %v", expected, actual)
	}
}

func TestIdentifierJSON(t *testing.T) {
	i := Identifier{
		ID:             "549300XMYB546ZI1F126",
		Type:           "Legal Entity Identifier",
		Category:       "GENERAL",
		CreationDate:   "2013-07-24T14:15:00Z",
		UpdatedDate:    "2023-08-04T17:33:45Z",
		ExpirationDate: "2020-01-16T00:32:00Z",
		Status:         "ACTIVE",
	}
	expected := `{"id":"549300XMYB546ZI1F126","id_type":"Legal Entity Identifier","category":"GENERAL","creation_date":"2013-07-24T14:15:00Z","update_date":"2023-08-04T17:33:45Z","expiration_date":"2020-01-16T00:32:00Z","status":"ACTIVE"}`
	actual, err := i.JSON()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(actual), expected) {
		t.Errorf("Expected JSON %v but got %v", expected, string(actual))
	}
}
