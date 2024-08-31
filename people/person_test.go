// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package people

import (
	"encoding/json"
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestPersonAssetType(t *testing.T) {
	p := Person{}
	if p.AssetType() != model.Person {
		t.Errorf("Expected asset type to be %v, but got %v", model.Person, p.AssetType())
	}
}

func TestPersonJSON(t *testing.T) {
	p := Person{
		FullName:   "John Doe",
		FirstName:  "John",
		MiddleName: "Jacob",
		FamilyName: "Doe",
	}

	expectedJSON := `{"full_name":"John Doe","first_name":"John","middle_name":"Jacob","family_name":"Doe"}`

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
