package fingerprint_test

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/owasp-amass/open-asset-model/fingerprint"
)

func TestFingerprint_AssetType(t *testing.T) {
	fp := fingerprint.Fingerprint{
		Value: "example",
		Type:  "example",
	}
	want := model.Fingerprint

	if got := fp.AssetType(); got != want {
		t.Errorf("Fingerprint.AssetType() = %v, want %v", got, want)
	}
}

func TestFingerprint_JSON(t *testing.T) {
	fp := fingerprint.Fingerprint{
		Value: "example",
		Type:  "example",
	}

	// Test AssetType method
	if fp.AssetType() != model.Fingerprint {
		t.Errorf("Expected asset type %s, but got %s", model.Fingerprint, fp.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"string":"example","type":"example"}`
	jsonData, err := fp.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(jsonData))
	}
}
