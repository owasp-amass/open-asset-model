// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package tls_certificate

import (
	"testing"
	"time"
)

func TestTLSCertificate(t *testing.T) {
	cert := TLSCertificate{
		SubjectCommonName: "www.example.org",
		IssuerCommonName:  "DigiCert TLS RSA SHA256 2020 CA1",
		NotBefore:         time.Date(2023, 1, 13, 12, 0, 0, 0, time.UTC),
		NotAfter:          time.Date(2024, 2, 13, 11, 59, 59, 0, time.UTC),
		FingerprintSHA1:   "F2AAD73D32683B716D2A7D61B51C6D5764AB3899",
		FingerprintSHA256: "5EF2F214260AB8F58E55EEA42E4AC04B0F171807D8D1185FDDD67470E9AB6096",
	}

	if cert.SubjectCommonName != "www.example.org" {
		t.Errorf("Failed to set the Subject Common Name field: %s", cert.SubjectCommonName)
	}

	if cert.IssuerCommonName != "DigiCert TLS RSA SHA256 2020 CA1" {
		t.Errorf("Failed to set the Issuer Common Name field: %s", cert.IssuerCommonName)
	}

	if cert.NotBefore != time.Date(2023, 1, 13, 12, 0, 0, 0, time.UTC) {
		t.Errorf("Failed to set the Not Before field: %s", cert.NotBefore)
	}

	if cert.NotAfter != time.Date(2024, 2, 13, 11, 59, 59, 0, time.UTC) {
		t.Errorf("Failed to set the Not After field: %s", cert.NotAfter)
	}

	if cert.FingerprintSHA1 != "F2AAD73D32683B716D2A7D61B51C6D5764AB3899" {
		t.Errorf("Failed to set the FingerprintSHA1 field: %s", cert.FingerprintSHA1)
	}

	if cert.FingerprintSHA256 != "5EF2F214260AB8F58E55EEA42E4AC04B0F171807D8D1185FDDD67470E9AB6096" {
		t.Errorf("Failed to set the FingerprintSHA256 field: %s", cert.FingerprintSHA256)
	}
}
