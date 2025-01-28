// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package open_asset_model

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
	FQDN             AssetType = "FQDN"
	Identifier       AssetType = "Identifier"
	IPAddress        AssetType = "IPAddress"
	IPNetRecord      AssetType = "IPNetRecord"
	Location         AssetType = "Location"
	Netblock         AssetType = "Netblock"
	Organization     AssetType = "Organization"
	Person           AssetType = "Person"
	Phone            AssetType = "Phone"
	Product          AssetType = "Product"
	ProductRelease   AssetType = "ProductRelease"
	Service          AssetType = "Service"
	TLSCertificate   AssetType = "TLSCertificate"
	URL              AssetType = "URL"
)

var AssetList = []AssetType{
	AutnumRecord, AutonomousSystem, ContactRecord, DomainRecord, EmailAddress, File, FQDN, Identifier,
	IPAddress, IPNetRecord, Location, Netblock, Organization, Person, Phone, Product, ProductRelease,
	Service, TLSCertificate, URL,
}
