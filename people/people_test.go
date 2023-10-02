package people_test

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/people"
)

func TestName_AssetType(t *testing.T) {
	n := Name{Name: "John Doe"}

	if n.AssetType() != model.Name {
		t.Errorf("Expected asset type %v but got %v", model.Name, n.AssetType())
	}
}

func TestName_JSON(t *testing.T) {
	n := Name{Name: "John Doe"}

	expectedJSON := `{"name":"John Doe"}`
	jsonBytes, err := n.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(jsonBytes, []byte(expectedJSON)) {
		t.Errorf("Expected JSON %v but got %v", expectedJSON, string(jsonBytes))
	}
}
