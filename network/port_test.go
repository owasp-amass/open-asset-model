package network_test

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/owasp-amass/open-asset-model/network"
)

func TestPort_AssetType(t *testing.T) {
	p := network.Port{}
	want := model.Port

	if got := p.AssetType(); got != want {
		t.Errorf("Port.AssetType() = %v, want %v", got, want)
	}
}

func TestPort_JSON(t *testing.T) {
	p := network.Port{
		Number:   80,
		Protocol: "tcp",
	}

	// Test AssetType method
	if p.AssetType() != model.Port {
		t.Errorf("Expected asset type %s, but got %s", model.Port, p.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"number":80,"protocol":"tcp"}`
	json, err := p.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
