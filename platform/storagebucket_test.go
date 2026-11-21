// Copyright © by Jeff Foley 2017-2026. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package platform

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestStorageBucketKey(t *testing.T) {
	want := "s3:example-public"
	b := StorageBucket{ID: want, Name: "example-public"}

	if got := b.Key(); got != want {
		t.Errorf("StorageBucket.Key() = %v, want %v", got, want)
	}
}

func TestStorageBucketAssetType(t *testing.T) {
	var _ model.Asset = StorageBucket{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*StorageBucket)(nil) // Verify the pointer properly implements the Asset interface

	b := StorageBucket{}
	want := model.StorageBucket

	if got := b.AssetType(); got != want {
		t.Errorf("StorageBucket.AssetType() = %v, want %v", got, want)
	}
}

func TestStorageBucketJSON(t *testing.T) {
	b := StorageBucket{
		ID:       "s3:example-public",
		Name:     "example-public",
		Provider: "Amazon Web Services",
		Service:  "s3",
		Exists:   true,
		Public:   true,
	}

	if b.AssetType() != model.StorageBucket {
		t.Errorf("Expected asset type %s, but got %s", model.StorageBucket, b.AssetType())
	}

	expectedJSON := `{"unique_id":"s3:example-public","bucket_name":"example-public","provider":"Amazon Web Services","service":"s3","exists":true,"public":true}`
	j, err := b.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(j) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(j))
	}
}
