// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"encoding/json"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/assert"
)

func TestPhoneKey(t *testing.T) {
	want := "12345556666"
	p := Phone{Raw: want}

	if got := p.Key(); got != want {
		t.Errorf("Phone.Key() = %v, want %v", got, want)
	}
}

func TestPhoneAssetType(t *testing.T) {
	var _ model.Asset = Phone{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*Phone)(nil) // Verify the pointer properly implements the  Asset interface.

	p := Phone{}
	assert.Equal(t, model.Phone, p.AssetType())
}

func TestPhoneJSON(t *testing.T) {
	p := Phone{
		Raw:           "123-456-7890 Ext. 123",
		E164:          "+1234567890",
		Type:          PhoneTypeMobile,
		CountryAbbrev: "US",
		CountryCode:   1,
		Ext:           "123",
	}

	expectedJSON := `{"raw":"123-456-7890 Ext. 123","e164":"+1234567890","type":"mobile","country_abbrev":"US","country_code":1,"ext":"123"}`

	jsonData, err := p.JSON()
	assert.NoError(t, err)

	var phoneJSON map[string]interface{}
	err = json.Unmarshal(jsonData, &phoneJSON)
	assert.NoError(t, err)

	var expected map[string]interface{}
	err = json.Unmarshal([]byte(expectedJSON), &expected)
	assert.NoError(t, err)

	assert.Equal(t, expected, phoneJSON)
}
