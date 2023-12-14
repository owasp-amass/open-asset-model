// tls_certificate.go

package tls_certificates

import (
	"encoding/json"
	"time"

	model "github.com/owasp-amass/open-asset-model"
)

// TLSCertificate represents a TLS Certificate asset
// Note: commented out fields may be introduced on next releases
type TLSCertificate struct {
	CommonName              string `json:"common_name"`
	SubjectOrganization     string `json:"subject_organization"`
	SubjectOrganizationUnit string `json:"subject_organization_unit"`
	SubjectStateOrProvince  string `json:"subject_state_or_province"`
	SubjectLocality         string `json:"subject_locality"`
	SubjectEmailAddress     string `json:"subject_email_address"`

	Issuer                 string `json:"issuer"`
	IssuerOrganization     string `json:"issuer_organization"`
	IssuerOrganizationUnit string `json:"issuer_organization_unit"`

	// - Validity start time
	NotBefore time.Time `json:"not_before"`
	// - Validity end time
	NotAfter time.Time `json:"not_after"`
	// - Subject Alternative Names
	SubjectAltNames []string `json:"subject_alt_names"`
	// - Signature Algorithm used in the cert
	SignatureAlgorithm string `json:"signature_algorithm"`
	// - Public Key Algorithm used in the cert
	PublicKeyAlgorithm string `json:"public_key_algorithm"`
	// - Public Key Size in bits
	FingerprintSHA1   string `json:"fingerprint_sha1"`   // - SHA-1 fingerprint of the certificate
	FingerprintSHA256 string `json:"fingerprint_sha256"` // - SHA-256 fingerprint of the certificate

	// - Serial Number of the certificate
	SerialNumber string `json:"serial_number"`
	// - Certificate version
	Version int `json:"version"`
	// - Whether this certificate is a Certificate Authority
	//IsCA bool `json:"is_ca"`
	// - Whether this certificate is self-signed
	//IsSelfSigned bool `json:"is_self_signed"`

	// - The purposes for which the public key can be used
	KeyUsage []string `json:"key_usage"`

	// - Purposes for which the certificate is used, in addition to or in place of Key Usage
	ExtendedKeyUsage []string `json:"extended_key_usage"`
	// - Points to CRL (Certificate Revocation List) distributor points
	CRLDistributionPoints []string `json:"crl_distribution_points"`
	// - URLs of the issuer
	IssuerURLs []string `json:"issuer_urls"`
	// - Online Certificate Status Protocol server URLs
	OCSPServer []string `json:"ocsp_server"`
	// - Certificate policies
	Policies []string `json:"policies"`
	// - The ID of the certificate's public key
	SubjectKeyID string `json:"subject_key_id"`
	// - The ID of the certificate's authority's public key
	AuthorityKeyID string `json:"authority_key_id"`

	// Added for later use:
	// SubjectInfo contains information about the subject of the certificate
	//SubjectInfo map[string]string `json:"subject_info"`

	// Embed more details if required, for example, details about Extensions
	//Extensions map[string]string `json:"extensions"` // To store other extensions in the certificate
}

// AssetType returns the asset type.
func (t TLSCertificate) AssetType() model.AssetType {
	return model.TLSCertificate
}

// JSON returns the JSON encoding of the struct.
func (t TLSCertificate) JSON() ([]byte, error) {
	return json.Marshal(t)
}
