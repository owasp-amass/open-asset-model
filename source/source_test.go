// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package source

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestSourceKey(t *testing.T) {
	want := "DNS"
	src := Source{Name: want}

	if got := src.Key(); got != want {
		t.Errorf("Source.Key() = %v, want %v", got, want)
	}
}

func TestSourceImplementsAsset(t *testing.T) {
	var _ model.Asset = Source{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*Source)(nil) // Verify the pointer properly implements the  Asset interface.
}

func TestSource(t *testing.T) {
	t.Run("Test successful creation of Source with valid name", func(t *testing.T) {
		src := Source{Name: "www.example.com", Confidence: 50}

		require.Equal(t, "www.example.com", src.Name)
		require.Equal(t, 50, src.Confidence)
		require.Equal(t, src.AssetType(), model.Source)
	})

	t.Run("Test successful JSON serialization of Source with valid name", func(t *testing.T) {
		src := Source{Name: "www.example.com", Confidence: 50}

		jsonData, err := src.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"name":"www.example.com","confidence":50}`, string(jsonData))
	})
}
