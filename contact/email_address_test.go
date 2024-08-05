package contact_test

import (
	"encoding/json"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/contact"
)

func TestEmailAddress_AssetType(t *testing.T) {
	e := EmailAddress{
		Address:   "test@example.com",
		LocalPart: "test",
		Domain:    "example.com",
	}

	if e.AssetType() != model.EmailAddress {
		t.Errorf("Expected asset type %s, got %s", model.EmailAddress, e.AssetType())
	}
}

func TestEmailAddress_JSON(t *testing.T) {
	e := EmailAddress{
		Address:   "test@example.com",
		LocalPart: "test",
		Domain:    "example.com",
	}

	expectedJSON := `{"address":"test@example.com","local_part":"test","domain":"example.com"}`

	jsonBytes, err := e.JSON()
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
	}

	if string(jsonBytes) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonBytes))
	}

	var unmarshaled EmailAddress
	err = json.Unmarshal(jsonBytes, &unmarshaled)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if unmarshaled != e {
		t.Errorf("Expected unmarshaled address %v, got %v", e, unmarshaled)
	}
}
