// Copyright © by Jeff Foley 2017-2026. All rights reserved.
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

var accountRels = map[string]map[RelationType][]AssetType{
	"id":   {SimpleRelation: {Identifier}},
	"user": {SimpleRelation: {Person, Organization}},
}

var autnumRecordRels = map[string]map[RelationType][]AssetType{
	"whois_server":      {SimpleRelation: {FQDN}},
	"registrant":        {SimpleRelation: {ContactRecord}},
	"admin_contact":     {SimpleRelation: {ContactRecord}},
	"abuse_contact":     {SimpleRelation: {ContactRecord}},
	"technical_contact": {SimpleRelation: {ContactRecord}},
	"rdap_url":          {SimpleRelation: {URL}},
}

var autonomousSystemRels = map[string]map[RelationType][]AssetType{
	"announces":    {SimpleRelation: {Netblock}},
	"registration": {SimpleRelation: {AutnumRecord}},
}

var contactRecordRels = map[string]map[RelationType][]AssetType{
	"fqdn":         {SimpleRelation: {FQDN}},
	"id":           {SimpleRelation: {Identifier}},
	"person":       {SimpleRelation: {Person}},
	"organization": {SimpleRelation: {Organization}},
	"location":     {SimpleRelation: {Location}},
	"phone":        {SimpleRelation: {Phone}},
	"url":          {SimpleRelation: {URL}},
}

var domainRecordRels = map[string]map[RelationType][]AssetType{
	"name_server":        {SimpleRelation: {FQDN}},
	"whois_server":       {SimpleRelation: {FQDN}},
	"registrar_contact":  {SimpleRelation: {ContactRecord}},
	"registrant_contact": {SimpleRelation: {ContactRecord}},
	"admin_contact":      {SimpleRelation: {ContactRecord}},
	"technical_contact":  {SimpleRelation: {ContactRecord}},
	"billing_contact":    {SimpleRelation: {ContactRecord}},
}

var fileRels = map[string]map[RelationType][]AssetType{
	"url":      {SimpleRelation: {URL}},
	"contains": {SimpleRelation: {ContactRecord, URL}},
	// The file's content hash as a shared Identifier (id_type sha256), so the
	// same document found at several URLs, or published by several parties,
	// is one hop apart, mirroring the Service and URL fingerprint relations.
	"fingerprint": {SimpleRelation: {Identifier}},
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
	"verified_for": {SimpleRelation: {Organization, Service}},
	// A domain's published mail policy names other domains: its SPF record
	// authorizes a domain's senders (include:/redirect=), and its DMARC record
	// sends reports to a domain (rua/ruf mailto). The raw policy stays in the
	// TXT record; these edges make the references pivotable.
	"spf_authorizes":   {SimpleRelation: {FQDN}},
	"dmarc_reports_to": {SimpleRelation: {FQDN}},
	// Identity-fabric relations: the domain's identity tenant (an Entra tenant
	// GUID or Okta org, as an Identifier) and the sibling domains that share the
	// same tenant federation trust.
	"identity_tenant": {SimpleRelation: {Identifier}},
	"federated_with":  {SimpleRelation: {FQDN}},
	// An object-storage bucket this name resolves to (an S3/GCS/Azure Blob
	// host), so a controlled name and the bucket it reaches are one hop apart.
	"bucket": {SimpleRelation: {StorageBucket}},
}

var fundsTransferRels = map[string]map[RelationType][]AssetType{
	"id":          {SimpleRelation: {Identifier}},
	"sender":      {SimpleRelation: {Account}},
	"recipient":   {SimpleRelation: {Account}},
	"third_party": {SimpleRelation: {Organization}},
}

var identifierRels = map[string]map[RelationType][]AssetType{
	"registration_agency": {SimpleRelation: {ContactRecord}},
	"issuing_authority":   {SimpleRelation: {ContactRecord}},
	"issuing_agent":       {SimpleRelation: {ContactRecord}},
}

var ipRels = map[string]map[RelationType][]AssetType{
	"port":       {PortRelation: {Service}},
	"ptr_record": {SimpleRelation: {FQDN}},
}

var ipnetRecordRels = map[string]map[RelationType][]AssetType{
	"whois_server":      {SimpleRelation: {FQDN}},
	"registrant":        {SimpleRelation: {ContactRecord}},
	"admin_contact":     {SimpleRelation: {ContactRecord}},
	"abuse_contact":     {SimpleRelation: {ContactRecord}},
	"technical_contact": {SimpleRelation: {ContactRecord}},
	"rdap_url":          {SimpleRelation: {URL}},
}

var locationRels = map[string]map[RelationType][]AssetType{
	"id": {SimpleRelation: {Identifier}},
}

var netblockRels = map[string]map[RelationType][]AssetType{
	"contains":     {SimpleRelation: {IPAddress}},
	"provider":     {SimpleRelation: {Organization}},
	"registration": {SimpleRelation: {IPNetRecord}},
}

var orgRels = map[string]map[RelationType][]AssetType{
	"id":                   {SimpleRelation: {Identifier}},
	"legal_address":        {SimpleRelation: {Location}},
	"hq_address":           {SimpleRelation: {Location}},
	"location":             {SimpleRelation: {Location}},
	"subsidiary":           {SimpleRelation: {Organization}},
	"org_unit":             {SimpleRelation: {Organization}},
	"account":              {SimpleRelation: {Account}},
	"member":               {SimpleRelation: {Person}},
	"website":              {SimpleRelation: {URL}},
	"social_media_profile": {SimpleRelation: {URL}},
	"funding_source":       {SimpleRelation: {Person, Organization}},
}

var personRels = map[string]map[RelationType][]AssetType{
	"id":      {SimpleRelation: {Identifier}},
	"address": {SimpleRelation: {Location}},
	"phone":   {SimpleRelation: {Phone}},
	"account": {SimpleRelation: {Account}},
}

var phoneRels = map[string]map[RelationType][]AssetType{
	"account": {SimpleRelation: {Account}},
	"contact": {SimpleRelation: {ContactRecord}},
}

var productRels = map[string]map[RelationType][]AssetType{
	"id":           {SimpleRelation: {Identifier}},
	"manufacturer": {SimpleRelation: {Organization}},
	"website":      {SimpleRelation: {URL}},
	"release":      {SimpleRelation: {ProductRelease}},
}

var productReleaseRels = map[string]map[RelationType][]AssetType{
	"id":      {SimpleRelation: {Identifier}},
	"website": {SimpleRelation: {URL}},
}

var serviceRels = map[string]map[RelationType][]AssetType{
	"provider":    {SimpleRelation: {Organization}},
	"certificate": {SimpleRelation: {TLSCertificate}},
	// A fingerprint of what the service served (favicon hash, JARM, page title,
	// content hash, analytics identifier, SSH host key) as a shared Identifier,
	// so services and hosts carrying the same value are one hop apart and can be
	// pivoted on in the graph. The Identifier's id_type names the fingerprint.
	"fingerprint":              {SimpleRelation: {Identifier}},
	"terms_of_service":         {SimpleRelation: {File, URL}},
	"product_used":             {SimpleRelation: {Product, ProductRelease}},
	"security_contact":         {SimpleRelation: {ContactRecord}},
	"security_policy":          {SimpleRelation: {URL}},
	"security_canonical":       {SimpleRelation: {URL}},
	"security_encryption":      {SimpleRelation: {URL}},
	"security_acknowledgments": {SimpleRelation: {URL}},
	"security_hiring":          {SimpleRelation: {URL}},
}

var storageBucketRels = map[string]map[RelationType][]AssetType{
	// The cloud provider that operates the bucket's namespace. Uses the same
	// "provider" convention as netblock.
	"provider": {SimpleRelation: {Organization}},
}

var tlscertRels = map[string]map[RelationType][]AssetType{
	"common_name":             {SimpleRelation: {FQDN}},
	"subject_contact":         {SimpleRelation: {ContactRecord}},
	"issuer_contact":          {SimpleRelation: {ContactRecord}},
	"san_dns_name":            {SimpleRelation: {FQDN}},
	"san_email_address":       {SimpleRelation: {Identifier}},
	"san_ip_address":          {SimpleRelation: {IPAddress}},
	"san_url":                 {SimpleRelation: {URL}},
	"issuing_certificate":     {SimpleRelation: {TLSCertificate}},
	"issuing_certificate_url": {SimpleRelation: {URL}},
	"ocsp_server":             {SimpleRelation: {URL}},
}

var urlRels = map[string]map[RelationType][]AssetType{
	"domain":     {SimpleRelation: {FQDN}},
	"ip_address": {SimpleRelation: {IPAddress}},
	"port":       {PortRelation: {Service}},
	"file":       {SimpleRelation: {File}},
	// A fingerprint (e.g. an analytics identifier) of the page at this URL, as a
	// shared Identifier, mirroring the service fingerprint relation.
	"fingerprint": {SimpleRelation: {Identifier}},
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
	case Account:
		relations = accountRels
	case AutnumRecord:
		relations = autnumRecordRels
	case AutonomousSystem:
		relations = autonomousSystemRels
	case ContactRecord:
		relations = contactRecordRels
	case DomainRecord:
		relations = domainRecordRels
	case File:
		relations = fileRels
	case FQDN:
		relations = fqdnRels
	case FundsTransfer:
		relations = fundsTransferRels
	case Identifier:
		relations = identifierRels
	case IPAddress:
		relations = ipRels
	case IPNetRecord:
		relations = ipnetRecordRels
	case Location:
		relations = locationRels
	case Netblock:
		relations = netblockRels
	case Organization:
		relations = orgRels
	case Person:
		relations = personRels
	case Phone:
		relations = phoneRels
	case Product:
		relations = productRels
	case ProductRelease:
		relations = productReleaseRels
	case Service:
		relations = serviceRels
	case StorageBucket:
		relations = storageBucketRels
	case TLSCertificate:
		relations = tlscertRels
	case URL:
		relations = urlRels
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
