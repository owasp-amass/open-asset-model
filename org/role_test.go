package org_test

import (
	"encoding/json"
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/org"
)

func TestRole_AssetType(t *testing.T) {
	r := Role{}
	if r.AssetType() != model.Role {
		t.Errorf("Expected asset type %v, but got %v", model.Role, r.AssetType())
	}
}

func TestRole_JSON(t *testing.T) {
	r := Role{Role: "admin"}
	expectedJSON := `{"role":"admin"}`

	jsonBytes, err := r.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(jsonBytes) != expectedJSON {
		t.Errorf("Expected JSON %v, but got %v", expectedJSON, string(jsonBytes))
	}
}

func TestRole_JSON_Unmarshal(t *testing.T) {
	jsonStr := `{"role":"admin"}`
	expectedRole := Role{Role: "admin"}

	var r Role
	err := json.Unmarshal([]byte(jsonStr), &r)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(r, expectedRole) {
		t.Errorf("Expected role %v, but got %v", expectedRole, r)
	}
}
