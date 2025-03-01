// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package general

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestPortRelationName(t *testing.T) {
	want := "port"
	pr := PortRelation{Name: want}

	if got := pr.Label(); got != want {
		t.Errorf("PortRelation.Label() = %v, want %v", got, want)
	}
}

func TestPortRelationImplementsRelation(t *testing.T) {
	var _ model.Relation = PortRelation{}       // Verify proper implementation of the Relation interface
	var _ model.Relation = (*PortRelation)(nil) // Verify *PortRelation properly implements the Relation interface.
}

func TestPortRelation(t *testing.T) {
	t.Run("Test successful creation of PortRelation with valid port number and protocol", func(t *testing.T) {
		pr := PortRelation{
			Name:       "port",
			PortNumber: 80,
			Protocol:   "tcp",
		}

		require.Equal(t, "port", pr.Name)
		require.Equal(t, 80, pr.PortNumber)
		require.Equal(t, "tcp", pr.Protocol)
		require.Equal(t, pr.RelationType(), model.PortRelation)
	})

	t.Run("Test successful JSON serialization of PortRelation with valid port number and protocol", func(t *testing.T) {
		pr := PortRelation{
			Name:       "port",
			PortNumber: 80,
			Protocol:   "tcp",
		}

		jsonData, err := pr.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"label":"port", "port_number":80, "protocol":"tcp"}`, string(jsonData))
	})
}
