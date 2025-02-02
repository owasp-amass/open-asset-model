// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package general

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestSimpleRelationName(t *testing.T) {
	want := "anything"
	sr := SimpleRelation{Name: want}

	if got := sr.Label(); got != want {
		t.Errorf("SimpleRelation.Label() = %v, want %v", got, want)
	}
}

func TestSimpleRelationImplementsRelation(t *testing.T) {
	var _ model.Relation = SimpleRelation{}       // Verify proper implementation of the Relation interface
	var _ model.Relation = (*SimpleRelation)(nil) // Verify *SimpleRelation properly implements the Relation interface.
}

func TestSimpleRelation(t *testing.T) {
	t.Run("Test successful creation of SimpleRelation", func(t *testing.T) {
		sr := SimpleRelation{Name: "anything"}

		require.Equal(t, "anything", sr.Name)
		require.Equal(t, sr.RelationType(), model.SimpleRelation)
	})

	t.Run("Test successful JSON serialization of SimpleRelation", func(t *testing.T) {
		sr := SimpleRelation{Name: "anything"}

		jsonData, err := sr.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"label":"anything"}`, string(jsonData))
	})
}

func TestSimplePropertyName(t *testing.T) {
	want := "anything"
	sp := SimpleProperty{
		PropertyName:  "anything",
		PropertyValue: "foobar",
	}

	if got := sp.Name(); got != want {
		t.Errorf("SimpleProperty.Name() = %v, want %v", got, want)
	}
}

func TestSimplePropertyValue(t *testing.T) {
	want := "foobar"
	sp := SimpleProperty{
		PropertyName:  "anything",
		PropertyValue: "foobar",
	}

	if got := sp.Value(); got != want {
		t.Errorf("SimpleProperty.Value() = %v, want %v", got, want)
	}
}

func TestSimplePropertyImplementsProperty(t *testing.T) {
	var _ model.Property = SimpleProperty{}       // Verify proper implementation of the Property interface
	var _ model.Property = (*SimpleProperty)(nil) // Verify *SimpleProperty properly implements the Property interface.
}

func TestSimpleProperty(t *testing.T) {
	t.Run("Test successful creation of SimpleProperty", func(t *testing.T) {
		sp := SimpleProperty{
			PropertyName:  "anything",
			PropertyValue: "foobar",
		}

		require.Equal(t, "anything", sp.PropertyName)
		require.Equal(t, "foobar", sp.PropertyValue)
		require.Equal(t, sp.PropertyType(), model.SimpleProperty)
	})

	t.Run("Test successful JSON serialization of SimpleProperty", func(t *testing.T) {
		sp := SimpleProperty{
			PropertyName:  "anything",
			PropertyValue: "foobar",
		}

		jsonData, err := sp.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"property_name":"anything", "property_value":"foobar"}`, string(jsonData))
	})
}
