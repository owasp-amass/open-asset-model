// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package account

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestAccountKey(t *testing.T) {
	want := "222333444"
	a := Account{
		ID:       want,
		Username: "test",
		Number:   "12345",
		Type:     "ACH",
	}

	if got := a.Key(); got != want {
		t.Errorf("Account.Key() = %v, want %v", got, want)
	}
}

func TestAccountAssetType(t *testing.T) {
	var _ model.Asset = Account{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*Account)(nil) // Verify the pointer properly implements the Asset interface.

	a := Account{}
	expected := model.Account
	actual := a.AssetType()

	if actual != expected {
		t.Errorf("Expected asset type %v but got %v", expected, actual)
	}
}

func TestAccountJSON(t *testing.T) {
	a := Account{
		ID:       "222333444",
		Type:     "ACH",
		Username: "test",
		Number:   "12345",
		Balance:  10000,
		Active:   true,
	}
	expected := `{"unique_id":"222333444","account_type":"ACH","username":"test","account_number":"12345","balance":10000,"active":true}`
	actual, err := a.JSON()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(actual), expected) {
		t.Errorf("Expected JSON %v but got %v", expected, string(actual))
	}
}
