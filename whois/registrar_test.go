// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package whois

import (
	"encoding/json"
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestRegistrarAssetType(t *testing.T) {
	reg := Registrar{}
	expected := model.Registrar
	actual := reg.AssetType()

	if actual != expected {
		t.Errorf("Expected asset type %v but got %v", expected, actual)
	}
}

func TestRegistrarJSON(t *testing.T) {
	reg := Registrar{
		Name:   "Registrar Name",
		IANAID: "12345",
	}
	expected := `{"name":"Registrar Name","iana_id":"12345"}`
	actual, err := reg.JSON()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(actual), expected) {
		t.Errorf("Expected JSON %v but got %v", expected, string(actual))
	}

	var reg2 Registrar
	err = json.Unmarshal(actual, &reg2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(reg, reg2) {
		t.Errorf("Expected registrar %v but got %v", reg, reg2)
	}
}
