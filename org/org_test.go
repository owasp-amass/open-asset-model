package org_test

import (
	"encoding/json"
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/org"
)

func TestOrganization_AssetType(t *testing.T) {
	org := Organization{}
	if org.AssetType() != model.Organization {
		t.Errorf("Expected asset type %v but got %v", model.Organization, org.AssetType())
	}
}

func TestOrganization_JSON(t *testing.T) {
	org := Organization{Organization: "TestOrg"}
	expectedJSON := `{"organization":"TestOrg"}`

	jsonBytes, err := org.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(jsonBytes) != expectedJSON {
		t.Errorf("Expected JSON %v but got %v", expectedJSON, string(jsonBytes))
	}
}

func TestOrganization_JSON_Unmarshal(t *testing.T) {
	jsonStr := `{"organization":"TestOrg"}`
	expectedOrg := Organization{Organization: "TestOrg"}

	var org Organization
	err := json.Unmarshal([]byte(jsonStr), &org)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(org, expectedOrg) {
		t.Errorf("Expected organization %v but got %v", expectedOrg, org)
	}
}
