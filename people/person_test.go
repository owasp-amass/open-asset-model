// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package people

import (
	"encoding/json"
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestPersonKey(t *testing.T) {
	want := "222333444"
	p := Person{
		ID:       want,
		FullName: "Jeff Foley",
	}

	if got := p.Key(); got != want {
		t.Errorf("Person.Key() = %v, want %v", got, want)
	}
}

func TestPersonAssetType(t *testing.T) {
	var _ model.Asset = Person{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*Person)(nil) // Verify the pointer properly implements the Asset interface.

	p := Person{}
	if p.AssetType() != model.Person {
		t.Errorf("Expected asset type to be %v, but got %v", model.Person, p.AssetType())
	}
}

func TestPersonJSON(t *testing.T) {
	p := Person{
		ID:         "222333444",
		FullName:   "John Jacob Doe",
		FirstName:  "John",
		MiddleName: "Jacob",
		FamilyName: "Doe",
		BirthDate:  "01/01/1970",
		Gender:     "Male",
	}

	expectedJSON := `{"unique_id":"222333444","full_name":"John Jacob Doe","first_name":"John","middle_name":"Jacob","family_name":"Doe","birth_date":"01/01/1970","gender":"Male"}`

	jsonData, err := p.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonData, &jsonMap)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedMap map[string]interface{}
	err = json.Unmarshal([]byte(expectedJSON), &expectedMap)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(jsonMap, expectedMap) {
		t.Errorf("Expected JSON to be %v, but got %v", expectedJSON, string(jsonData))
	}
}
