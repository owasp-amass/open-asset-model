// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package open_asset_model

import (
	"strings"
)

type Asset interface {
	Key() string
	AssetType() AssetType
	JSON() ([]byte, error)
}

type AssetType string

const (
	AutnumRecord     AssetType = "AutnumRecord"
	AutonomousSystem AssetType = "AutonomousSystem"
	ContactRecord    AssetType = "ContactRecord"
	DomainRecord     AssetType = "DomainRecord"
	EmailAddress     AssetType = "EmailAddress"
	File             AssetType = "File"
	Fingerprint      AssetType = "Fingerprint"
	FQDN             AssetType = "FQDN"
	IPAddress        AssetType = "IPAddress"
	IPNetRecord      AssetType = "IPNetRecord"
	Location         AssetType = "Location"
	Netblock         AssetType = "Netblock"
	NetworkEndpoint  AssetType = "NetworkEndpoint"
	Organization     AssetType = "Organization"
	Person           AssetType = "Person"
	Phone            AssetType = "Phone"
	Service          AssetType = "Service"
	SocketAddress    AssetType = "SocketAddress"
	Source           AssetType = "Source"
	TLSCertificate   AssetType = "TLSCertificate"
	URL              AssetType = "URL"
)

var AssetList = []AssetType{
	AutnumRecord, AutonomousSystem, ContactRecord, DomainRecord, EmailAddress, File,
	Fingerprint, FQDN, IPAddress, IPNetRecord, Location, Netblock, NetworkEndpoint,
	Organization, Person, Phone, Service, SocketAddress, Source, TLSCertificate, URL,
}

var locationRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
}

var phoneRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
}

var emailRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"domain":       {FQDN},
}

var domainRecordRels = map[string][]AssetType{
	"source":             {Source},
	"monitored_by":       {Source},
	"name_server":        {FQDN},
	"whois_server":       {FQDN},
	"registrar_contact":  {ContactRecord},
	"registrant_contact": {ContactRecord},
	"admin_contact":      {ContactRecord},
	"technical_contact":  {ContactRecord},
	"billing_contact":    {ContactRecord},
	"associated_with":    {AutnumRecord, DomainRecord, IPNetRecord},
}

var autnumRecordRels = map[string][]AssetType{
	"source":            {Source},
	"monitored_by":      {Source},
	"whois_server":      {FQDN},
	"registrant":        {ContactRecord},
	"admin_contact":     {ContactRecord},
	"abuse_contact":     {ContactRecord},
	"technical_contact": {ContactRecord},
	"rdap_url":          {URL},
	"associated_with":   {AutnumRecord, DomainRecord, IPNetRecord},
}

var ipnetRecordRels = map[string][]AssetType{
	"source":            {Source},
	"monitored_by":      {Source},
	"whois_server":      {FQDN},
	"registrant":        {ContactRecord},
	"admin_contact":     {ContactRecord},
	"abuse_contact":     {ContactRecord},
	"technical_contact": {ContactRecord},
	"rdap_url":          {URL},
	"associated_with":   {AutnumRecord, DomainRecord, IPNetRecord},
}

var personRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
}

var orgRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
}

var ipRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"port":         {SocketAddress},
	"ptr_record":   {FQDN},
}

var netblockRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"contains":     {IPAddress},
	"registration": {IPNetRecord},
}

var autonomousSystemRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"announces":    {Netblock},
	"registration": {AutnumRecord},
}

var fileRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"url":          {URL},
	"contains":     {ContactRecord, URL},
}

var fqdnRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"port":         {NetworkEndpoint},
	"a_record":     {IPAddress},
	"aaaa_record":  {IPAddress},
	"cname_record": {FQDN},
	"ns_record":    {FQDN},
	"ptr_record":   {FQDN},
	"mx_record":    {FQDN},
	"srv_record":   {FQDN, IPAddress},
	"node":         {FQDN},
	"registration": {DomainRecord},
}

var tlscertRels = map[string][]AssetType{
	"source":                  {Source},
	"monitored_by":            {Source},
	"common_name":             {FQDN},
	"subject_contact":         {ContactRecord},
	"issuer_contact":          {ContactRecord},
	"san_dns_name":            {FQDN},
	"san_email_address":       {EmailAddress},
	"san_ip_address":          {IPAddress},
	"san_url":                 {URL},
	"issuing_certificate":     {TLSCertificate},
	"issuing_certificate_url": {URL},
	"ocsp_server":             {URL},
	"associated_with":         {AutnumRecord, DomainRecord, IPNetRecord},
}

var socketAddressRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"service":      {Service},
}

var networkEndpointRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"service":      {Service},
}

var urlRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"domain":       {FQDN},
	"ip_address":   {IPAddress},
	"port":         {SocketAddress, NetworkEndpoint},
	"service":      {Service},
	"file":         {File},
}

var fingerprintRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
}

var contactRecordRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"person":       {Person},
	"organization": {Organization},
	"location":     {Location},
	"email":        {EmailAddress},
	"phone":        {Phone},
	"url":          {URL},
}

var sourceRels = map[string][]AssetType{
	"contact_record": {ContactRecord},
}

var serviceRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"fingerprint":  {Fingerprint},
	"certificate":  {TLSCertificate},
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

// GetAssetOutgoingRelations returns the relation types allowed to be used
// when the subject is the asset type provided in the parameter.
// Providing an invalid subject causes a return value of nil.
func GetTransformAssetTypes(subject AssetType, relation string) []AssetType {
	relations := assetTypeRelations(subject)
	if relations == nil {
		return nil
	}

	rtype := strings.ToLower(relation)
	if atypes, ok := relations[rtype]; ok {
		return atypes
	}
	return nil
}

func assetTypeRelations(atype AssetType) map[string][]AssetType {
	var relations map[string][]AssetType

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
	case NetworkEndpoint:
		relations = networkEndpointRels
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
	case SocketAddress:
		relations = socketAddressRels
	case URL:
		relations = urlRels
	case Fingerprint:
		relations = fingerprintRels
	case ContactRecord:
		relations = contactRecordRels
	case Source:
		relations = sourceRels
	case Service:
		relations = serviceRels
	default:
		return nil
	}

	return relations
}

// ValidRelationship returns true if the relation is valid in the taxonomy
// when outgoing from the source asset type to the destination asset type.
func ValidRelationship(src AssetType, relation string, destination AssetType) bool {
	atypes := GetTransformAssetTypes(src, relation)
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
