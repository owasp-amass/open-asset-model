// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package open_asset_model

import (
	"strings"
)

type Relation interface {
	Label() string
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

var locationRels = map[string]map[RelationType][]AssetType{}

var phoneRels = map[string]map[RelationType][]AssetType{}

var emailRels = map[string]map[RelationType][]AssetType{
	"domain": {SimpleRelation: {FQDN}},
}

var domainRecordRels = map[string]map[RelationType][]AssetType{
	"name_server":        {SimpleRelation: {FQDN}},
	"whois_server":       {SimpleRelation: {FQDN}},
	"registrar_contact":  {SimpleRelation: {ContactRecord}},
	"registrant_contact": {SimpleRelation: {ContactRecord}},
	"admin_contact":      {SimpleRelation: {ContactRecord}},
	"technical_contact":  {SimpleRelation: {ContactRecord}},
	"billing_contact":    {SimpleRelation: {ContactRecord}},
	"associated_with":    {SimpleRelation: {AutnumRecord, DomainRecord, IPNetRecord}},
}

var autnumRecordRels = map[string]map[RelationType][]AssetType{
	"whois_server":      {SimpleRelation: {FQDN}},
	"registrant":        {SimpleRelation: {ContactRecord}},
	"admin_contact":     {SimpleRelation: {ContactRecord}},
	"abuse_contact":     {SimpleRelation: {ContactRecord}},
	"technical_contact": {SimpleRelation: {ContactRecord}},
	"rdap_url":          {SimpleRelation: {URL}},
	"associated_with":   {SimpleRelation: {AutnumRecord, DomainRecord, IPNetRecord}},
}

var ipnetRecordRels = map[string]map[RelationType][]AssetType{
	"whois_server":      {SimpleRelation: {FQDN}},
	"registrant":        {SimpleRelation: {ContactRecord}},
	"admin_contact":     {SimpleRelation: {ContactRecord}},
	"abuse_contact":     {SimpleRelation: {ContactRecord}},
	"technical_contact": {SimpleRelation: {ContactRecord}},
	"rdap_url":          {SimpleRelation: {URL}},
	"associated_with":   {SimpleRelation: {AutnumRecord, DomainRecord, IPNetRecord}},
}

var personRels = map[string]map[RelationType][]AssetType{}

var orgRels = map[string]map[RelationType][]AssetType{}

var ipRels = map[string]map[RelationType][]AssetType{
	"port":       {PortRelation: {Service}},
	"ptr_record": {SimpleRelation: {FQDN}},
}

var netblockRels = map[string]map[RelationType][]AssetType{
	"contains":     {SimpleRelation: {IPAddress}},
	"registration": {SimpleRelation: {IPNetRecord}},
}

var autonomousSystemRels = map[string]map[RelationType][]AssetType{
	"announces":    {SimpleRelation: {Netblock}},
	"registration": {SimpleRelation: {AutnumRecord}},
}

var fileRels = map[string]map[RelationType][]AssetType{
	"url":      {SimpleRelation: {URL}},
	"contains": {SimpleRelation: {ContactRecord, URL}},
}

var fqdnRels = map[string]map[RelationType][]AssetType{
	"port": {PortRelation: {Service}},
	"dns_record": {
		BasicDNSRelation: {FQDN, IPAddress},
		PrefDNSRelation:  {FQDN},
		SRVDNSRelation:   {FQDN},
	},
	"node":         {SimpleRelation: {FQDN}},
	"registration": {SimpleRelation: {DomainRecord}},
}

var tlscertRels = map[string]map[RelationType][]AssetType{
	"common_name":             {SimpleRelation: {FQDN}},
	"subject_contact":         {SimpleRelation: {ContactRecord}},
	"issuer_contact":          {SimpleRelation: {ContactRecord}},
	"san_dns_name":            {SimpleRelation: {FQDN}},
	"san_email_address":       {SimpleRelation: {EmailAddress}},
	"san_ip_address":          {SimpleRelation: {IPAddress}},
	"san_url":                 {SimpleRelation: {URL}},
	"issuing_certificate":     {SimpleRelation: {TLSCertificate}},
	"issuing_certificate_url": {SimpleRelation: {URL}},
	"ocsp_server":             {SimpleRelation: {URL}},
	"associated_with":         {SimpleRelation: {AutnumRecord, DomainRecord, IPNetRecord}},
}

var urlRels = map[string]map[RelationType][]AssetType{
	"domain":     {SimpleRelation: {FQDN}},
	"ip_address": {SimpleRelation: {IPAddress}},
	"port":       {PortRelation: {Service}},
	"file":       {SimpleRelation: {File}},
}

var fingerprintRels = map[string]map[RelationType][]AssetType{}

var contactRecordRels = map[string]map[RelationType][]AssetType{
	"person":       {SimpleRelation: {Person}},
	"organization": {SimpleRelation: {Organization}},
	"location":     {SimpleRelation: {Location}},
	"email":        {SimpleRelation: {EmailAddress}},
	"phone":        {SimpleRelation: {Phone}},
	"url":          {SimpleRelation: {URL}},
}

var serviceRels = map[string]map[RelationType][]AssetType{
	"fingerprint": {SimpleRelation: {Fingerprint}},
	"certificate": {SimpleRelation: {TLSCertificate}},
}

// GetAssetOutgoingRelations returns the relation types allowed to be used
// when the subject is the asset type provided in the parameter.
// Providing an invalid subject causes a return value of nil.
func GetAssetOutgoingRelations(subject AssetType) []string {
	relations := assetTypeRelations(subject)
	if relations == nil {
		return nil
	}

	var rtypes []string
	for k := range relations {
		rtypes = append(rtypes, k)
	}
	return rtypes
}

// GetTransformAssetTypes returns the asset types allowed to be assigned
// when the subject is the asset type provided in the parameter, along
// with the provided label and RelationType.
// Providing an invalid subject causes a return value of nil.
func GetTransformAssetTypes(subject AssetType, label string, rtype RelationType) []AssetType {
	relations := assetTypeRelations(subject)
	if relations == nil {
		return nil
	}

	var results []AssetType
	label = strings.ToLower(label)
	m := make(map[AssetType]struct{})
	for r, atypes := range relations[label] {
		if r != rtype {
			continue
		}
		for _, t := range atypes {
			if _, found := m[t]; !found {
				m[t] = struct{}{}
				results = append(results, t)
			}
		}
	}
	return results
}

func assetTypeRelations(atype AssetType) map[string]map[RelationType][]AssetType {
	var relations map[string]map[RelationType][]AssetType

	switch atype {
	case IPAddress:
		relations = ipRels
	case Netblock:
		relations = netblockRels
	case AutonomousSystem:
		relations = autonomousSystemRels
	case File:
		relations = fileRels
	case FQDN:
		relations = fqdnRels
	case DomainRecord:
		relations = domainRecordRels
	case AutnumRecord:
		relations = autnumRecordRels
	case IPNetRecord:
		relations = ipnetRecordRels
	case Location:
		relations = locationRels
	case Phone:
		relations = phoneRels
	case EmailAddress:
		relations = emailRels
	case Person:
		relations = personRels
	case Organization:
		relations = orgRels
	case TLSCertificate:
		relations = tlscertRels
	case URL:
		relations = urlRels
	case Fingerprint:
		relations = fingerprintRels
	case ContactRecord:
		relations = contactRecordRels
	case Service:
		relations = serviceRels
	default:
		return nil
	}

	return relations
}

// ValidRelationship returns true if the relation is valid in the taxonomy
// when outgoing from the source asset type to the destination asset type.
func ValidRelationship(src AssetType, label string, rtype RelationType, destination AssetType) bool {
	atypes := GetTransformAssetTypes(src, label, rtype)
	if atypes == nil {
		return false
	}

	for _, atype := range atypes {
		if atype == destination {
			return true
		}
	}
	return false
}
