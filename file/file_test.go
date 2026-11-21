// Copyright © by Jeff Foley 2017-2026. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package file

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestFileKey(t *testing.T) {
	want := "file:///var/html/index.html"
	f := File{URL: want, Name: "index.html", Type: "text/html"}

	if got := f.Key(); got != want {
		t.Errorf("File.Key() = %v, want %v", got, want)
	}
}

func TestFileAssetType(t *testing.T) {
	var _ model.Asset = File{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*File)(nil) // Verify the pointer properly implements the  Asset interface.

	f := File{}
	expected := model.File
	actual := f.AssetType()

	if actual != expected {
		t.Errorf("Expected asset type %v but got %v", expected, actual)
	}
}

func TestFileJSON(t *testing.T) {
	f := File{
		URL:  "https://example.com/reports/annual.pdf",
		Name: "annual.pdf",
		Type: "application/pdf",
		Size: 482133,
	}
	expected := `{"url":"https://example.com/reports/annual.pdf","name":"annual.pdf","type":"application/pdf","size":482133}`
	actual, err := f.JSON()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(actual), expected) {
		t.Errorf("Expected JSON %v but got %v", expected, string(actual))
	}
}

func TestFileJSONOmitsUnknownSize(t *testing.T) {
	f := File{URL: "https://example.com/a.docx", Type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document"}
	got, err := f.JSON()
	if err != nil {
		t.Fatal(err)
	}

	want := `{"url":"https://example.com/a.docx","type":"application/vnd.openxmlformats-officedocument.wordprocessingml.document"}`
	if string(got) != want {
		t.Errorf("JSON = %s, want %s", got, want)
	}
}
