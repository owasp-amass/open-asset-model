package fingerprint

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Fingerprint represents a fingerprint.
type Fingerprint struct {
	Value string `json:"value"` // Fingerprint string
	Type  string `json:"type"`  // Fingerprint type
}

// AssetType returns the asset type.
func (f Fingerprint) AssetType() model.AssetType {
	return model.Fingerprint
}

// JSON returns the JSON encoding of the struct.
func (f Fingerprint) JSON() ([]byte, error) {
	return json.Marshal(f)
}
