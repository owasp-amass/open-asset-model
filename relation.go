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
	BasicDNSRelation RelationType = "BasicDNSRelation"
	PortRelation     RelationType = "PortRelation"
	PrefDNSRelation  RelationType = "PrefDNSRelation"
	SimpleRelation   RelationType = "SimpleRelation"
	SRVDNSRelation   RelationType = "SRVDNSRelation"
)

var RelationList = []RelationType{
	BasicDNSRelation, PortRelation, PrefDNSRelation, SimpleRelation, SRVDNSRelation,
}
