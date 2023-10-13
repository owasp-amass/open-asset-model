package tls_certificates_test

import (
	"fmt"
	"os"
	"time"

	. "github.com/owasp-amass/open-asset-model/tls_certificates"
)

func ExampleTLSCertificate() {

	TLSCertificate := TLSCertificate{
		CommonName:          "www.example.org",
		SubjectOrganization: "Internet Corporation for Assigned Name and Numbers",

		Issuer:             "DigiCert TLS RSA SHA256 2020 CA1",
		IssuerOrganization: "DigiCert Inc",

		NotBefore: time.Date(2023, 1, 13, 12, 0, 0, 0, time.UTC),
		NotAfter:  time.Date(2024, 2, 13, 11, 59, 59, 0, time.UTC),

		FingerprintSHA1:   "F2AAD73D32683B716D2A7D61B51C6D5764AB3899",
		FingerprintSHA256: "5EF2F214260AB8F58E55EEA42E4AC04B0F171807D8D1185FDDD67470E9AB6096",
	}

	rval := TLSCertificate.CommonName == "www.example.org"
	if !rval {
		fmt.Println(TLSCertificate.CommonName)
		os.Exit(1)
	} else {
		rval = TLSCertificate.SubjectOrganization == "Internet Corporation for Assigned Name and Numbers"
	}
	if !rval {
		fmt.Println(TLSCertificate.SubjectOrganization)
		os.Exit(1)
	} else {
		rval = TLSCertificate.Issuer == "DigiCert TLS RSA SHA256 2020 CA1"
	}
	if !rval {
		fmt.Println(TLSCertificate.Issuer)
		os.Exit(1)
	} else {
		rval = TLSCertificate.IssuerOrganization == "DigiCert Inc"
	}
	if !rval {
		fmt.Println(TLSCertificate.IssuerOrganization)
		os.Exit(1)
	} else {
		rval = TLSCertificate.NotBefore == time.Date(2023, 1, 13, 12, 0, 0, 0, time.UTC)
	}
	if !rval {
		fmt.Println(TLSCertificate.NotBefore)
		os.Exit(1)
	} else {
		rval = TLSCertificate.NotAfter == time.Date(2024, 2, 13, 11, 59, 59, 0, time.UTC)
	}
	if !rval {
		fmt.Println(TLSCertificate.NotAfter)
		os.Exit(1)
	} else {
		rval = TLSCertificate.FingerprintSHA1 == "F2AAD73D32683B716D2A7D61B51C6D5764AB3899"
	}
	if !rval {
		fmt.Println(TLSCertificate.FingerprintSHA1)
		os.Exit(1)
	} else {
		rval = TLSCertificate.FingerprintSHA256 == "5EF2F214260AB8F58E55EEA42E4AC04B0F171807D8D1185FDDD67470E9AB6096"
	}
	if !rval {
		fmt.Println(TLSCertificate.FingerprintSHA256)
		os.Exit(1)
	}

	// Output:
	// true
}
