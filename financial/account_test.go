// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package financial

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestAccountKey(t *testing.T) {
	want := "ACH:12345"
	a := Account{
		Number: "12345",
		Type:   "ACH",
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
		Number:             "12345",
		Type:               "ACH",
		DomesticBankNumber: "222333444",
		Balance:            10000,
		Currency:           "US",
		InterestRate:       0.05,
		PercentageRate:     0.05,
		FeesAndCommissions: 100,
		Active:             true,
		CountryCode:        "US",
	}
	expected := `{"account_number":"12345","account_type":"ACH","domestic_bank_number":"222333444","balance":10000,"currency":"US","interest_rate":0.05,"percentage_rate":0.05,"fees_and_commissions":100,"active":true,"country_code":"US"}`
	actual, err := a.JSON()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(actual), expected) {
		t.Errorf("Expected JSON %v but got %v", expected, string(actual))
	}
}
