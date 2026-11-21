// Copyright © by Jeff Foley 2017-2026. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package contact

import (
	"encoding/json"
	"fmt"

	model "github.com/owasp-amass/open-asset-model"
)

// GeoLocation represents a geographic position: a point on the Earth given by
// its latitude and longitude in decimal degrees (WGS 84), as observed, with or
// without a street address. A photo's GPS, an IP geolocation record, a map
// feature and a registry's coordinates are each a GeoLocation; a street address
// geocoded to a point links to one.
//
// Its identity is the position rounded to six decimal places (about 11 cm), so
// observations of the same point converge on one asset. Nearby points are not
// merged: clustering them is a judgment made when the data is read. Altitude
// and Uncertainty describe the observation and are not part of the identity;
// Uncertainty, in meters, matters as much as the position itself, since a
// photo's GPS is good to meters and an IP geolocation often only to a city.
//
// It should support relationships such as the following:
// - A street address at the position (StreetAddress --geolocation--> GeoLocation)
// - An IP address geolocated to it (IPAddress --geolocation--> GeoLocation)
// - A file that records it (File --contains--> GeoLocation)
type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	// Altitude is meters above the WGS 84 ellipsoid, when known.
	Altitude float64 `json:"altitude,omitempty"`
	// Uncertainty is the radius in meters within which the position lies,
	// when known (RFC 5870's "u" parameter).
	Uncertainty float64 `json:"uncertainty,omitempty"`
}

// Key implements the Asset interface: the position rounded to six decimal
// places, "latitude,longitude".
func (g GeoLocation) Key() string {
	return fmt.Sprintf("%.6f,%.6f", g.Latitude, g.Longitude)
}

// AssetType implements the Asset interface.
func (g GeoLocation) AssetType() model.AssetType {
	return model.GeoLocation
}

// JSON implements the Asset interface.
func (g GeoLocation) JSON() ([]byte, error) {
	return json.Marshal(g)
}

// Valid reports whether the position is a point on the Earth.
func (g GeoLocation) Valid() bool {
	return g.Latitude >= -90 && g.Latitude <= 90 && g.Longitude >= -180 && g.Longitude <= 180
}
