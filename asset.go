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
	WHOISRecord    AssetType = "WHOISRecord"
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
)

var AssetList = []AssetType{
	IPAddress, Netblock, ASN, RIROrg, FQDN, WHOISRecord, Location,
	Phone, EmailAddress, Person, Organization, Registrar, Registrant,
	SocketAddress, URL, Fingerprint, TLSCertificate,
}

var locationRels = map[string][]AssetType{}

var phoneRels = map[string][]AssetType{}

var emailRels = map[string][]AssetType{}

var whoisRels = map[string][]AssetType{
	"published_by": {Registrar},
	"name_server":  {FQDN},
	"reseller":     {Organization},

	"admin_org":      {Organization},
	"admin_person":   {Person},
	"admin_phone":    {Phone},
	"admin_email":    {EmailAddress},
	"admin_location": {Location},

	"tech_org":      {Organization},
	"tech_person":   {Person},
	"tech_phone":    {Phone},
	"tech_email":    {EmailAddress},
	"tech_location": {Location},

	"billing_org":      {Organization},
	"billing_person":   {Person},
	"billing_phone":    {Phone},
	"billing_email":    {EmailAddress},
	"billing_location": {Location},

	"registrant_org":      {Organization},
	"registrant_person":   {Person},
	"registrant_phone":    {Phone},
	"registrant_email":    {EmailAddress},
	"registrant_location": {Location},
}

var personRels = map[string][]AssetType{
	"phone_number": {Phone},
	"email":        {EmailAddress},
	"location":     {Location},
}

var orgRels = map[string][]AssetType{
	"rir_org":      {RIROrg},
	"location":     {Location},
	"phone_number": {Phone},
	"email":        {EmailAddress},
	"operates":     {Registrar},
}

var registrarRels = map[string][]AssetType{
	"abuse_email":  {EmailAddress},
	"abuse_phone":  {Phone},
	"whois_server": {FQDN},
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
	"registration": {WHOISRecord},
}

var tlscertRels = map[string][]AssetType{
	"common_name":               {FQDN},
	"subject_organization":      {Organization},
	"subject_organization_unit": {Organization},
	"subject_state_or_province": {Location},
	"subject_locality":          {Location},
	"subject_email":             {EmailAddress},
	"issuer":                    {FQDN},
	"issuer_organization":       {Organization},
	"issuer_organization_unit":  {Organization},
	"subject_alt_names":         {FQDN},
	"issuer_urls":               {URL},
	"ocsp_server":               {URL},
	"jarm":                      {Fingerprint},
}

var socketAddressRels = map[string][]AssetType{}

var urlRels = map[string][]AssetType{
	"port":       {SocketAddress},
	"domain":     {FQDN},
	"ip_address": {IPAddress},
}

var fingerprintRels = map[string][]AssetType{}

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
	case WHOISRecord:
		relations = whoisRels
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
