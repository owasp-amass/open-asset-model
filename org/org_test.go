package org_test

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/org"
)

func TestOrganization_AssetType(t *testing.T) {
	org := Organization{}
	expected := model.Organization
	actual := org.AssetType()

	if actual != expected {
		t.Errorf("Expected asset type %v but got %v", expected, actual)
	}
}

func TestOrganization_JSON(t *testing.T) {
	org := Organization{
		OrgName:  "Acme Inc.",
		Industry: "Technology",
	}
	expected := `{"org_name":"Acme Inc.","industry":"Technology"}`
	actual, err := org.JSON()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(actual), expected) {
		t.Errorf("Expected JSON %v but got %v", expected, string(actual))
	}
}
