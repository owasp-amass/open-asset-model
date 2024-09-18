// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package certificate

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestTLSCertificateKey(t *testing.T) {
	want := "12345"
	cert := TLSCertificate{SerialNumber: want}

	if got := cert.Key(); got != want {
		t.Errorf("TLSCertificate.Key() = %v, want %v", got, want)
	}
}

func TestTLSCertificateAssetType(t *testing.T) {
	var _ model.Asset = TLSCertificate{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*TLSCertificate)(nil) // Verify the pointer properly implements the  Asset interface.

	cert := TLSCertificate{}
	want := model.TLSCertificate

	if got := cert.AssetType(); got != want {
		t.Errorf("TLSCertificate.AssetType() = %v, want %v", got, want)
	}
}

func TestTLSCertificateJSON(t *testing.T) {
	cert := TLSCertificate{
		SubjectCommonName: "www.example.org",
		IssuerCommonName:  "DigiCert TLS RSA SHA256 2020 CA1",
		NotBefore:         "2006-01-02T15:04:05Z07:00",
		NotAfter:          "2006-01-02T15:04:05Z07:00",
	}

	// test AssetType method
	if cert.AssetType() != model.TLSCertificate {
		t.Errorf("Expected asset type %s, but got %s", model.TLSCertificate, cert.AssetType())
	}

	// test JSON method
	expectedJSON := `{"version":"","serial_number":"","subject_common_name":"www.example.org","issuer_common_name":"DigiCert TLS RSA SHA256 2020 CA1","not_before":"2006-01-02T15:04:05Z07:00","not_after":"2006-01-02T15:04:05Z07:00","key_usage":null,"ext_key_usage":null,"signature_algorithm":"","public_key_algorithm":"","is_ca":false,"crl_distribution_points":null,"subject_key_id":"","authority_key_id":""}`
	json, err := cert.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
