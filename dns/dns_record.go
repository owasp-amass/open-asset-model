// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

type RRHeader struct {
	RRType int `json:"rr_type"`
	Class  int `json:"class"`
	TTL    int `json:"ttl"`
}

// BasicDNSRelation is a relation in the graph representing a basic DNS resource record.
type BasicDNSRelation struct {
	Name   string   `json:"label"`
	Header RRHeader `json:"header"`
}

// Label implements the Relation interface.
func (r BasicDNSRelation) Label() string {
	return r.Name
}

// RelationType implements the Relation interface.
func (r BasicDNSRelation) RelationType() model.RelationType {
	return model.BasicDNSRelation
}

// JSON implements the Relation interface.
func (r BasicDNSRelation) JSON() ([]byte, error) {
	return json.Marshal(r)
}

// PrefDNSRelation is a relation in the graph representing a DNS resource record with preference information.
type PrefDNSRelation struct {
	Name       string   `json:"label"`
	Header     RRHeader `json:"header"`
	Preference int      `json:"preference"`
}

// Label implements the Relation interface.
func (r PrefDNSRelation) Label() string {
	return r.Name
}

// RelationType implements the Relation interface.
func (r PrefDNSRelation) RelationType() model.RelationType {
	return model.PrefDNSRelation
}

// JSON implements the Relation interface.
func (r PrefDNSRelation) JSON() ([]byte, error) {
	return json.Marshal(r)
}

// SRVDNSRelation is a relation in the graph representing a DNS SRV resource record.
type SRVDNSRelation struct {
	Name     string   `json:"label"`
	Header   RRHeader `json:"header"`
	Priority int      `json:"priority"`
	Weight   int      `json:"weight"`
	Port     int      `json:"port"`
}

// Label implements the Relation interface.
func (r SRVDNSRelation) Label() string {
	return r.Name
}

// RelationType implements the Relation interface.
func (r SRVDNSRelation) RelationType() model.RelationType {
	return model.SRVDNSRelation
}

// JSON implements the Relation interface.
func (r SRVDNSRelation) JSON() ([]byte, error) {
	return json.Marshal(r)
}

// DNSRecordProperty represents a DNS resource record that does not refer to another asset in the graph.
type DNSRecordProperty struct {
	PropertyName string   `json:"property_name"`
	Header       RRHeader `json:"header"`
	Data         string   `json:"data"`
}

// Name implements the Property interface.
func (r DNSRecordProperty) Name() string {
	return r.PropertyName
}

// Value implements the Property interface.
func (r DNSRecordProperty) Value() string {
	return r.Data
}

// PropertyType implements the Property interface.
func (r DNSRecordProperty) PropertyType() model.PropertyType {
	return model.DNSRecordProperty
}

// JSON implements the Property interface.
func (r DNSRecordProperty) JSON() ([]byte, error) {
	return json.Marshal(r)
}
