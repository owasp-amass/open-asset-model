// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package open_asset_model

type Asset interface {
	AssetType() AssetType
	JSON() ([]byte, error)
}

type AssetType string

const (
	IPAddress        AssetType = "IPAddress"
	Netblock         AssetType = "Netblock"
	AutonomousSystem AssetType = "AutonomousSystem"
	FQDN             AssetType = "FQDN"
	DomainRecord     AssetType = "DomainRecord"
	AutnumRecord     AssetType = "AutnumRecord"
	Location         AssetType = "Location"
	Phone            AssetType = "Phone"
	EmailAddress     AssetType = "EmailAddress"
	Person           AssetType = "Person"
	Organization     AssetType = "Organization"
	SocketAddress    AssetType = "SocketAddress"
	URL              AssetType = "URL"
	Fingerprint      AssetType = "Fingerprint"
	TLSCertificate   AssetType = "TLSCertificate"
	ContactRecord    AssetType = "ContactRecord"
	Source           AssetType = "Source"
)

var AssetList = []AssetType{
	IPAddress, Netblock, AutonomousSystem, FQDN, DomainRecord, AutnumRecord,
	Location, Phone, EmailAddress, Person, Organization, SocketAddress, URL,
	Fingerprint, TLSCertificate, ContactRecord, Source,
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
}

var netblockRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"contains":     {IPAddress},
}

var autonomousSystemRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"announces":    {Netblock},
	"registration": {AutnumRecord},
}

var fqdnRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
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
	"subject_alt_names":       {FQDN},
	"san_dns_name":            {FQDN},
	"san_email_address":       {EmailAddress},
	"san_ip_address":          {IPAddress},
	"san_url":                 {URL},
	"issuing_certificate_url": {URL},
	"ocsp_server":             {URL},
}

var socketAddressRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
}

var urlRels = map[string][]AssetType{
	"source":       {Source},
	"monitored_by": {Source},
	"port":         {SocketAddress},
	"domain":       {FQDN},
	"ip_address":   {IPAddress},
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

// ValidRelationship returns true if the relation is valid in the taxonomy
// when outgoing from the source asset type to the destination asset type.
func ValidRelationship(src AssetType, relation string, destination AssetType) bool {
	var relations map[string][]AssetType

	switch src {
	case IPAddress:
		relations = ipRels
	case Netblock:
		relations = netblockRels
	case AutonomousSystem:
		relations = autonomousSystemRels
	case FQDN:
		relations = fqdnRels
	case DomainRecord:
		relations = domainRecordRels
	case AutnumRecord:
		relations = autnumRecordRels
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
	default:
		return false
	}

	if atypes, ok := relations[relation]; ok {
		for _, atype := range atypes {
			if atype == destination {
				return true
			}
		}
	}
	return false
}
