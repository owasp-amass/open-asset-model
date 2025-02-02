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
	Account          AssetType = "Account"
	AutnumRecord     AssetType = "AutnumRecord"
	AutonomousSystem AssetType = "AutonomousSystem"
	ContactRecord    AssetType = "ContactRecord"
	DomainRecord     AssetType = "DomainRecord"
	File             AssetType = "File"
	FQDN             AssetType = "FQDN"
	FundsTransfer    AssetType = "FundsTransfer"
	Identifier       AssetType = "Identifier"
	IPAddress        AssetType = "IPAddress"
	IPNetRecord      AssetType = "IPNetRecord"
	Institution      AssetType = "Institution"
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
	Account, AutnumRecord, AutonomousSystem, ContactRecord, DomainRecord, File, FQDN, FundsTransfer,
	Identifier, IPAddress, IPNetRecord, Location, Netblock, Organization, Person, Phone, Product,
	ProductRelease, Service, TLSCertificate, URL,
}
