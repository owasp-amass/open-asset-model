// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package property

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

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
