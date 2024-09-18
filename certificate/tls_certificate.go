// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package certificate

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

const (
	KeyUsageDigitalSignature  string = "Digital Signature"
	KeyUsageContentCommitment string = "Content Commitment"
	KeyUsageKeyEncipherment   string = "Key Encipherment"
	KeyUsageDataEncipherment  string = "Data Encipherment"
	KeyUsageKeyAgreement      string = "Key Agreement"
	KeyUsageCertSign          string = "Certificate Sign"
	KeyUsageCRLSign           string = "CRL Sign"
	KeyUsageEncipherOnly      string = "Encipher Only"
	KeyUsageDecipherOnly      string = "Decipher Only"

	ExtKeyUsageAny                            string = "Any Usage"
	ExtKeyUsageServerAuth                     string = "TLS Server Authentication"
	ExtKeyUsageClientAuth                     string = "TLS Client Authentication"
	ExtKeyUsageCodeSigning                    string = "Code Signing"
	ExtKeyUsageEmailProtection                string = "E-mail Protection"
	ExtKeyUsageIPSECEndSystem                 string = "IPSec End System"
	ExtKeyUsageIPSECTunnel                    string = "IPSec Tunnel"
	ExtKeyUsageIPSECUser                      string = "IPSec User"
	ExtKeyUsageTimeStamping                   string = "Trusted Timestamping"
	ExtKeyUsageOCSPSigning                    string = "OCSP Signing"
	ExtKeyUsageMicrosoftServerGatedCrypto     string = "Microsoft Server Gated Crypto"
	ExtKeyUsageNetscapeServerGatedCrypto      string = "Netscape Server Gated Crypto"
	ExtKeyUsageMicrosoftCommercialCodeSigning string = "Microsoft Commercial Code Signing"
	ExtKeyUsageMicrosoftKernelCodeSigning     string = "Microsoft Kernel Code Signing"
)

// TLSCertificate represents a TLS Certificate asset.
type TLSCertificate struct {
	Version               string   `json:"version"`
	SerialNumber          string   `json:"serial_number"`
	SubjectCommonName     string   `json:"subject_common_name"`
	IssuerCommonName      string   `json:"issuer_common_name"`
	NotBefore             string   `json:"not_before"`
	NotAfter              string   `json:"not_after"`
	KeyUsage              []string `json:"key_usage"`
	ExtKeyUsage           []string `json:"ext_key_usage"`
	SignatureAlgorithm    string   `json:"signature_algorithm"`
	PublicKeyAlgorithm    string   `json:"public_key_algorithm"`
	IsCA                  bool     `json:"is_ca"`
	CRLDistributionPoints []string `json:"crl_distribution_points"`
	SubjectKeyID          string   `json:"subject_key_id"`
	AuthorityKeyID        string   `json:"authority_key_id"`
}

// Key implements the Asset interface.
func (t TLSCertificate) Key() string {
	return t.SerialNumber
}

// AssetType implements the Asset interface.
func (t TLSCertificate) AssetType() model.AssetType {
	return model.TLSCertificate
}

// JSON implements the Asset interface.
func (t TLSCertificate) JSON() ([]byte, error) {
	return json.Marshal(t)
}
