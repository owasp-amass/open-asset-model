// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package relation

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestSimpleRelationImplementsRelation(t *testing.T) {
	var _ model.Relation = SimpleRelation{}       // Verify proper implementation of the Relation interface
	var _ model.Relation = (*SimpleRelation)(nil) // Verify *SimpleRelation properly implements the Relation interface.
}

func TestSimpleRelation(t *testing.T) {
	t.Run("Test successful creation of SimpleRelation", func(t *testing.T) {
		sr := SimpleRelation{}

		require.Equal(t, sr.RelationType(), model.SimpleRelation)
	})

	t.Run("Test successful JSON serialization of SimpleRelation", func(t *testing.T) {
		sr := SimpleRelation{}

		jsonData, err := sr.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{}`, string(jsonData))
	})
}
