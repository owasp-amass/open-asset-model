// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package property

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestSourcePropertyName(t *testing.T) {
	want := "anything"
	sp := SourceProperty{
		Source:     "anything",
		Confidence: 80,
	}

	if got := sp.Name(); got != want {
		t.Errorf("SourceProperty.Name() = %v, want %v", got, want)
	}
}

func TestSourcePropertyValue(t *testing.T) {
	want := "80"
	sp := SourceProperty{
		Source:     "anything",
		Confidence: 80,
	}

	if got := sp.Value(); got != want {
		t.Errorf("SourceProperty.Value() = %v, want %v", got, want)
	}
}

func TestSourcePropertyImplementsProperty(t *testing.T) {
	var _ model.Property = SourceProperty{}       // Verify proper implementation of the Property interface
	var _ model.Property = (*SourceProperty)(nil) // Verify *SourceProperty properly implements the Property interface.
}

func TestSourceProperty(t *testing.T) {
	t.Run("Test successful creation of SourceProperty", func(t *testing.T) {
		sp := SourceProperty{
			Source:     "anything",
			Confidence: 80,
		}

		require.Equal(t, "anything", sp.Source)
		require.Equal(t, 80, sp.Confidence)
		require.Equal(t, sp.PropertyType(), model.SourceProperty)
	})

	t.Run("Test successful JSON serialization of SourceProperty", func(t *testing.T) {
		sp := SourceProperty{
			Source:     "anything",
			Confidence: 80,
		}

		jsonData, err := sp.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"name":"anything", "confidence":80}`, string(jsonData))
	})
}
