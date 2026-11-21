// Copyright © by Jeff Foley 2017-2026. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"encoding/json"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestGeoLocationKey(t *testing.T) {
	g := GeoLocation{Latitude: 42.468639734, Longitude: -75.064, Uncertainty: 5}
	if got := g.Key(); got != "42.468640,-75.064000" {
		t.Errorf("Key() = %s", got)
	}
	// the observation's accuracy is not part of the identity
	if g.Key() != (GeoLocation{Latitude: 42.4686397, Longitude: -75.064, Uncertainty: 50000}).Key() {
		t.Error("the same point with another uncertainty has another key")
	}
}

func TestGeoLocationAssetType(t *testing.T) {
	var _ model.Asset = GeoLocation{}
	var _ model.Asset = (*GeoLocation)(nil)

	if (GeoLocation{}).AssetType() != model.GeoLocation {
		t.Errorf("asset type %s", (GeoLocation{}).AssetType())
	}
}

func TestGeoLocationJSON(t *testing.T) {
	g := GeoLocation{Latitude: 42.46864, Longitude: -75.064, Altitude: 350.5, Uncertainty: 10}
	data, err := g.JSON()
	if err != nil {
		t.Fatal(err)
	}
	want := `{"latitude":42.46864,"longitude":-75.064,"altitude":350.5,"uncertainty":10}`
	if string(data) != want {
		t.Errorf("JSON %s, want %s", data, want)
	}

	var back GeoLocation
	if err := json.Unmarshal(data, &back); err != nil || back != g {
		t.Errorf("round trip %+v, %v", back, err)
	}

	// the equator and the prime meridian are positions, so zero is written
	if data, _ := (GeoLocation{}).JSON(); string(data) != `{"latitude":0,"longitude":0}` {
		t.Errorf("zero position JSON %s", data)
	}
}

func TestGeoLocationValid(t *testing.T) {
	for _, g := range []GeoLocation{{Latitude: 90, Longitude: 180}, {Latitude: -90, Longitude: -180}, {}} {
		if !g.Valid() {
			t.Errorf("%+v is a position", g)
		}
	}
	for _, g := range []GeoLocation{{Latitude: 91}, {Longitude: -181}} {
		if g.Valid() {
			t.Errorf("%+v is not a position", g)
		}
	}
}
