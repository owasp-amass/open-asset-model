package url_test

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/owasp-amass/open-asset-model/url"
)

func TestURL_AssetType(t *testing.T) {
	u := url.URL{}
	want := model.URL

	if got := u.AssetType(); got != want {
		t.Errorf("URL.AssetType() = %v, want %v", got, want)
	}
}

func TestURL_JSON(t *testing.T) {
	u := url.URL{
		Raw:      "http://user:pass@example.com:8080/path?option1=value1&option2=value2#fragment",
		Scheme:   "http",
		Username: "user",
		Password: "pass",
		Host:     "example.com",
		Port:     8080,
		Path:     "/path",
		Options:  "option1=value1&option2=value2",
		Fragment: "fragment",
	}

	// Test AssetType method
	if u.AssetType() != model.URL {
		t.Errorf("Expected asset type %s, but got %s", model.URL, u.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"url":"http://user:pass@example.com:8080/path?option1=value1&option2=value2#fragment","scheme":"http","username":"user","password":"pass","host":"example.com","port":8080,"path":"/path","options":"option1=value1&option2=value2","fragment":"fragment"}`
	json, err := u.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	t.Log(string(json))
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
