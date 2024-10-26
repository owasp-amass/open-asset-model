// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package open_asset_model

type Relation interface {
	RelationType() RelationType
	JSON() ([]byte, error)
}

type RelationType string

const (
	SimpleRelation RelationType = "SimpleRelation"
	PortRelation   RelationType = "PortRelation"
)

var RelationList = []RelationType{
	PortRelation,
	SimpleRelation,
}
