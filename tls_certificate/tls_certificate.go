// Copyright © by Jeff Foley 2023-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package tls_certificate

import (
	"encoding/json"
	"time"

	model "github.com/owasp-amass/open-asset-model"
)

// TLSCertificate represents a TLS Certificate asset
type TLSCertificate struct {
	// certificate version
	Version int `json:"version"`
	// serial Number of the certificate
	SerialNumber string `json:"serial_number"`

	SubjectCommonName         string `json:"subject_common_name"`
	SubjectCountry            string `json:"subject_country"`
	SubjectOrganization       string `json:"subject_organization"`
	SubjectOrganizationalUnit string `json:"subject_organizational_unit"`
	SubjectLocality           string `json:"subject_locality"`
	SubjectProvince           string `json:"subject_province"`
	SubjectStreetAddress      string `json:"subject_street_address"`
	SubjectPostalCode         string `json:"subject_postal_code"`

	IssuerCommonName         string `json:"issuer_common_name"`
	IssuerCountry            string `json:"issuer_country"`
	IssuerOrganization       string `json:"issuer_organization"`
	IssuerOrganizationalUnit string `json:"issuer_organizational_unit"`
	IssuerLocality           string `json:"issuer_locality"`
	IssuerProvince           string `json:"issuer_province"`
	IssuerStreetAddress      string `json:"issuer_street_address"`
	IssuerPostalCode         string `json:"issuer_postal_code"`

	// validity start time
	NotBefore time.Time `json:"not_before"`
	// validity end time
	NotAfter time.Time `json:"not_after"`
	// the purposes for which the public key can be used
	KeyUsage []string `json:"key_usage"`

	// Subject Alternative Names
	DNSNames       []string `json:"san_dns_names"`
	EmailAddresses []string `json:"san_email_addresses"`
	IPAddresses    []string `json:"san_ip_addresses"`
	URIs           []string `json:"san_uris"`

	// Signature Algorithm used in the cert
	SignatureAlgorithm string `json:"signature_algorithm"`
	// Public Key Algorithm used in the cert
	PublicKeyAlgorithm string `json:"public_key_algorithm"`
	FingerprintSHA1    string `json:"fingerprint_sha1"`   // - SHA-1 fingerprint of the certificate
	FingerprintSHA256  string `json:"fingerprint_sha256"` // - SHA-256 fingerprint of the certificate

	// whether this certificate is a Certificate Authority
	IsCA bool `json:"is_ca"`
	// points to CRL (Certificate Revocation List) distributor points
	CRLDistributionPoints []string `json:"crl_distribution_points"`
	// URLs of the issuer
	IssuingCertificateURL []string `json:"issuing_certificate_urls"`
	// Online Certificate Status Protocol server URLs
	OCSPServer []string `json:"ocsp_server"`
	// the ID of the certificate's public key
	SubjectKeyID string `json:"subject_key_id"`
	// the ID of the certificate's authority's public key
	AuthorityKeyID string `json:"authority_key_id"`
}

// AssetType returns the asset type.
func (t TLSCertificate) AssetType() model.AssetType {
	return model.TLSCertificate
}

// JSON returns the JSON encoding of the struct.
func (t TLSCertificate) JSON() ([]byte, error) {
	return json.Marshal(t)
}
