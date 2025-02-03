// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package financial

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestFundsTransferKey(t *testing.T) {
	want := "222333444"
	ft := FundsTransfer{ID: want, Amount: 42.0}

	if got := ft.Key(); got != want {
		t.Errorf("FundsTransfer.Key() = %v, want %v", got, want)
	}
}

func TestFundsTransferAssetType(t *testing.T) {
	var _ model.Asset = FundsTransfer{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*FundsTransfer)(nil) // Verify the pointer properly implements the Asset interface.

	ft := FundsTransfer{}
	expected := model.FundsTransfer
	actual := ft.AssetType()

	if actual != expected {
		t.Errorf("Expected asset type %v but got %v", expected, actual)
	}
}

func TestFundsTransferJSON(t *testing.T) {
	ft := FundsTransfer{
		ID:              "222333444",
		Amount:          42.0,
		ReferenceNumber: "555666777",
		Currency:        "US",
		Method:          "ACH",
		ExchangeDate:    "2013-07-24T14:15:00Z",
		ExchangeRate:    0.9,
	}
	expected := `{"unique_id":"222333444","amount":42,"reference_number":"555666777","currency":"US","transfer_method":"ACH","exchange_date":"2013-07-24T14:15:00Z","exchange_rate":0.9}`
	actual, err := ft.JSON()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(actual), expected) {
		t.Errorf("Expected JSON %v but got %v", expected, string(actual))
	}
}
