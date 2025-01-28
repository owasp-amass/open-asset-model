// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package platform

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestProductKey(t *testing.T) {
	want := "12345"
	p := Product{
		UniqueID: want,
		Name:     "OWASP Amass",
		Type:     "Information Security",
	}

	if got := p.Key(); got != want {
		t.Errorf("Product.Key() = %v, want %v", got, want)
	}
}

func TestProductAssetType(t *testing.T) {
	var _ model.Asset = Product{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*Product)(nil) // Verify the pointer properly implements the  Asset interface.

	p := Product{}
	expected := model.Product
	actual := p.AssetType()

	if actual != expected {
		t.Errorf("Expected asset type %v but got %v", expected, actual)
	}
}

func TestProductJSON(t *testing.T) {
	p := Product{
		UniqueID:        "12345",
		Name:            "OWASP Amass",
		Type:            "Attack Surface Management",
		Category:        "Information Security",
		Description:     "In-depth attack surface mapping and asset discovery",
		CountryOfOrigin: "US",
	}
	expected := `{"unique_id":"12345","name":"OWASP Amass","product_type":"Attack Surface Management","category":"Information Security","description":"In-depth attack surface mapping and asset discovery","country_of_origin":"US"}`
	actual, err := p.JSON()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(actual), expected) {
		t.Errorf("Expected JSON %v but got %v", expected, string(actual))
	}
}

func TestProductReleaseKey(t *testing.T) {
	want := "12345"
	p := ProductRelease{
		UniqueID:      want,
		VersionNumber: "v4.2.0",
	}

	if got := p.Key(); got != want {
		t.Errorf("ProductRelease.Key() = %v, want %v", got, want)
	}
}

func TestProductReleaseAssetType(t *testing.T) {
	var _ model.Asset = ProductRelease{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*ProductRelease)(nil) // Verify the pointer properly implements the  Asset interface.

	p := ProductRelease{}
	expected := model.ProductRelease
	actual := p.AssetType()

	if actual != expected {
		t.Errorf("Expected asset type %v but got %v", expected, actual)
	}
}

func TestProductReleaseJSON(t *testing.T) {
	p := ProductRelease{
		UniqueID:      "12345",
		VersionNumber: "v4.2.0",
		ModelNumber:   "N/A",
		SerialNumber:  "N/A",
		ReleaseDate:   "2023-09-10T14:15:00Z",
	}
	expected := `{"unique_id":"12345","version_number":"v4.2.0","model_number":"N/A","serial_number":"N/A","release_date":"2023-09-10T14:15:00Z"}`
	actual, err := p.JSON()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(actual), expected) {
		t.Errorf("Expected JSON %v but got %v", expected, string(actual))
	}
}
