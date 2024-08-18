// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package open_asset_model

type Asset interface {
	AssetType() AssetType
	JSON() ([]byte, error)
}

type AssetType string

const (
	IPAddress      AssetType = "IPAddress"
	Netblock       AssetType = "Netblock"
	ASN            AssetType = "ASN"
	RIROrg         AssetType = "RIROrg"
	FQDN           AssetType = "FQDN"
	DomainRecord   AssetType = "DomainRecord"
	Location       AssetType = "Location"
	Phone          AssetType = "Phone"
	EmailAddress   AssetType = "EmailAddress"
	Person         AssetType = "Person"
	Organization   AssetType = "Organization"
	Registrar      AssetType = "Registrar"
	Registrant     AssetType = "Registrant"
	SocketAddress  AssetType = "SocketAddress"
	URL            AssetType = "URL"
	Fingerprint    AssetType = "Fingerprint"
	TLSCertificate AssetType = "TLSCertificate"
	ContactRecord  AssetType = "ContactRecord"
)

var AssetList = []AssetType{
	IPAddress, Netblock, ASN, RIROrg, FQDN, DomainRecord, Location,
	Phone, EmailAddress, Person, Organization, Registrar, Registrant,
	SocketAddress, URL, Fingerprint, TLSCertificate, ContactRecord,
}

var locationRels = map[string][]AssetType{}

var phoneRels = map[string][]AssetType{}

var emailRels = map[string][]AssetType{}

var domainRecordRels = map[string][]AssetType{
	"published_by":       {Registrar},
	"name_server":        {FQDN},
	"reseller":           {Organization},
	"registrar_contact":  {ContactRecord},
	"registrant_contact": {ContactRecord},
	"admin_contact":      {ContactRecord},
	"technical_contact":  {ContactRecord},
	"billing_contact":    {ContactRecord},
}

var personRels = map[string][]AssetType{}

var orgRels = map[string][]AssetType{
	"rir_org":        {RIROrg},
	"contact_record": {ContactRecord},
	"operates":       {Registrar},
}

var registrarRels = map[string][]AssetType{
	"contact_record": {ContactRecord},
	"abuse_contact":  {ContactRecord},
	"whois_server":   {FQDN},
}

var ipRels = map[string][]AssetType{
	"port": {SocketAddress},
}

var netblockRels = map[string][]AssetType{
	"contains": {IPAddress},
}

var asnRels = map[string][]AssetType{
	"announces":  {Netblock},
	"managed_by": {RIROrg},
}

var rirOrgRels = map[string][]AssetType{}

var fqdnRels = map[string][]AssetType{
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
	"jarm":                    {Fingerprint},
}

var socketAddressRels = map[string][]AssetType{}

var urlRels = map[string][]AssetType{
	"port":       {SocketAddress},
	"domain":     {FQDN},
	"ip_address": {IPAddress},
}

var fingerprintRels = map[string][]AssetType{}

var contactRecordRels = map[string][]AssetType{
	"person":       {Person},
	"organization": {Organization},
	"location":     {Location},
	"email":        {EmailAddress},
	"phone":        {Phone},
	"url":          {URL},
}

// ValidRelationship returns true if the relation is valid in the taxonomy
// when outgoing from the source asset type to the destination asset type.
func ValidRelationship(source AssetType, relation string, destination AssetType) bool {
	var relations map[string][]AssetType

	switch source {
	case IPAddress:
		relations = ipRels
	case Netblock:
		relations = netblockRels
	case ASN:
		relations = asnRels
	case RIROrg:
		relations = rirOrgRels
	case FQDN:
		relations = fqdnRels
	case DomainRecord:
		relations = domainRecordRels
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
	case Registrar:
		relations = registrarRels
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
