// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package open_asset_model

type Property interface {
	Name() string
	Value() string
	PropertyType() PropertyType
	JSON() ([]byte, error)
}

type PropertyType string

const (
	SimpleProperty PropertyType = "SimpleProperty"
	VulnProperty   PropertyType = "VulnProperty"
)

var PropertyList = []PropertyType{
	SimpleProperty, VulnProperty,
}
