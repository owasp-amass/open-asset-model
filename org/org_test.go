// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package org

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestOrganization_AssetType(t *testing.T) {
	o := Organization{}
	expected := model.Organization
	actual := o.AssetType()

	if actual != expected {
		t.Errorf("Expected asset type %v but got %v", expected, actual)
	}
}

func TestOrganization_JSON(t *testing.T) {
	o := Organization{
		Name:     "Acme Inc.",
		Industry: "Technology",
	}
	expected := `{"name":"Acme Inc.","industry":"Technology"}`
	actual, err := o.JSON()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(actual), expected) {
		t.Errorf("Expected JSON %v but got %v", expected, string(actual))
	}
}
