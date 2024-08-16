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
	Version               int       `json:"version"`
	SerialNumber          string    `json:"serial_number"`
	SubjectCommonName     string    `json:"subject_common_name"`
	IssuerCommonName      string    `json:"issuer_common_name"`
	NotBefore             time.Time `json:"not_before"`
	NotAfter              time.Time `json:"not_after"`
	KeyUsage              []string  `json:"key_usage"`
	SignatureAlgorithm    string    `json:"signature_algorithm"`
	PublicKeyAlgorithm    string    `json:"public_key_algorithm"`
	FingerprintSHA1       string    `json:"fingerprint_sha1"`   // - SHA-1 fingerprint of the certificate
	FingerprintSHA256     string    `json:"fingerprint_sha256"` // - SHA-256 fingerprint of the certificate
	IsCA                  bool      `json:"is_ca"`
	CRLDistributionPoints []string  `json:"crl_distribution_points"`
	SubjectKeyID          string    `json:"subject_key_id"`
	AuthorityKeyID        string    `json:"authority_key_id"`
}

// AssetType returns the asset type.
func (t TLSCertificate) AssetType() model.AssetType {
	return model.TLSCertificate
}

// JSON returns the JSON encoding of the struct.
func (t TLSCertificate) JSON() ([]byte, error) {
	return json.Marshal(t)
}
