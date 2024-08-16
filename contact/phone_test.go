// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"encoding/json"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/assert"
)

func TestPhone_AssetType(t *testing.T) {
	p := Phone{}
	assert.Equal(t, model.Phone, p.AssetType())
}

func TestPhone_JSON(t *testing.T) {
	p := Phone{
		Raw:              "123-456-7890",
		E164:             "+1234567890",
		Type:             "mobile",
		CountryAbbrev:    "US",
		CountryCode:      1,
		SubscriberNumber: "2345678901",
		Ext:              "123",
	}

	expectedJSON := `{"raw":"123-456-7890","e164":"+1234567890","type":"mobile","country_abbrev":"US","country_code":1,"subscriber_number":"2345678901","ext":"123"}`

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
