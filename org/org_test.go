// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package org

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestOrganizationKey(t *testing.T) {
	want := "222333444"
	o := Organization{
		ID:   want,
		Name: "OWASP Foundation",
	}

	if got := o.Key(); got != want {
		t.Errorf("Organization.Key() = %v, want %v", got, want)
	}
}

func TestOrganizationAssetType(t *testing.T) {
	var _ model.Asset = Organization{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*Organization)(nil) // Verify the pointer properly implements the Asset interface.

	o := Organization{}
	expected := model.Organization
	actual := o.AssetType()

	if actual != expected {
		t.Errorf("Expected asset type %v but got %v", expected, actual)
	}
}

func TestOrganizationJSON(t *testing.T) {
	o := Organization{
		ID:             "222333444",
		Name:           "Acme Inc.",
		LegalName:      "Acme Inc.",
		FoundingDate:   "2013-07-24T14:15:00Z",
		Jurisdiction:   "US-DE",
		RegistrationID: "123456789",
		Industry:       "Technology",
		Active:         true,
		NonProfit:      false,
		NumOfEmployees: 10000,
	}
	expected := `{"unique_id":"222333444","name":"Acme Inc.","legal_name":"Acme Inc.","founding_date":"2013-07-24T14:15:00Z","jurisdiction":"US-DE","registration_id":"123456789","industry":"Technology","active":true,"num_of_employees":10000}`
	actual, err := o.JSON()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(actual), expected) {
		t.Errorf("Expected JSON %v but got %v", expected, string(actual))
	}
}
