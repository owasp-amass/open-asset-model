// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package relation

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestBasicDNSRelationName(t *testing.T) {
	want := "dns_record"
	br := BasicDNSRelation{Name: want}

	if got := br.Label(); got != want {
		t.Errorf("BasicDNSRelation.Label() = %v, want %v", got, want)
	}
}

func TestBasicDNSRelationImplementsRelation(t *testing.T) {
	var _ model.Relation = BasicDNSRelation{}       // Verify proper implementation of the Relation interface
	var _ model.Relation = (*BasicDNSRelation)(nil) // Verify *BasicDNSRelation properly implements the Relation interface.
}

func TestBasicDNSRelation(t *testing.T) {
	t.Run("Test successful creation of BasicDNSRelation with valid resource record header", func(t *testing.T) {
		dr := BasicDNSRelation{
			Name: "dns_record",
			Header: RRHeader{
				RRType: 1,
				Class:  1,
				TTL:    86400,
			},
		}

		require.Equal(t, "dns_record", dr.Name)
		require.Equal(t, 1, dr.Header.RRType)
		require.Equal(t, 1, dr.Header.Class)
		require.Equal(t, 86400, dr.Header.TTL)
		require.Equal(t, dr.RelationType(), model.BasicDNSRelation)
	})

	t.Run("Test successful JSON serialization of BasicDNSRelation with valid resource record header", func(t *testing.T) {
		dr := BasicDNSRelation{
			Name: "dns_record",
			Header: RRHeader{
				RRType: 1,
				Class:  1,
				TTL:    86400,
			},
		}

		jsonData, err := dr.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"label":"dns_record", "header":{"rr_type":1, "class":1, "ttl":86400}}`, string(jsonData))
	})
}

func TestPrefDNSRelationName(t *testing.T) {
	want := "dns_record"
	br := PrefDNSRelation{Name: want}

	if got := br.Label(); got != want {
		t.Errorf("PrefDNSRelation.Label() = %v, want %v", got, want)
	}
}

func TestPrefDNSRelationImplementsRelation(t *testing.T) {
	var _ model.Relation = PrefDNSRelation{}       // Verify proper implementation of the Relation interface
	var _ model.Relation = (*PrefDNSRelation)(nil) // Verify *PrefDNSRelation properly implements the Relation interface.
}

func TestPrefDNSRelation(t *testing.T) {
	t.Run("Test successful creation of PrefDNSRelation with valid resource record header and preference", func(t *testing.T) {
		pr := PrefDNSRelation{
			Name: "dns_record",
			Header: RRHeader{
				RRType: 1,
				Class:  1,
				TTL:    86400,
			},
			Preference: 5,
		}

		require.Equal(t, "dns_record", pr.Name)
		require.Equal(t, 1, pr.Header.RRType)
		require.Equal(t, 1, pr.Header.Class)
		require.Equal(t, 86400, pr.Header.TTL)
		require.Equal(t, 5, pr.Preference)
		require.Equal(t, pr.RelationType(), model.PrefDNSRelation)
	})

	t.Run("Test successful JSON serialization of PrefDNSRelation with valid resource record header and preference", func(t *testing.T) {
		pr := PrefDNSRelation{
			Name: "dns_record",
			Header: RRHeader{
				RRType: 1,
				Class:  1,
				TTL:    86400,
			},
			Preference: 5,
		}

		jsonData, err := pr.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"label":"dns_record", "header":{"rr_type":1, "class":1, "ttl":86400}, "preference":5}`, string(jsonData))
	})
}

func TestSRVDNSRelationName(t *testing.T) {
	want := "dns_record"
	br := SRVDNSRelation{Name: want}

	if got := br.Label(); got != want {
		t.Errorf("SRVDNSRelation.Label() = %v, want %v", got, want)
	}
}

func TestSRVDNSRelationImplementsRelation(t *testing.T) {
	var _ model.Relation = SRVDNSRelation{}       // Verify proper implementation of the Relation interface
	var _ model.Relation = (*SRVDNSRelation)(nil) // Verify *SRVDNSRelation properly implements the Relation interface.
}

func TestSRVDNSRelation(t *testing.T) {
	t.Run("Test successful creation of SRVDNSRelation", func(t *testing.T) {
		sr := SRVDNSRelation{
			Name: "dns_record",
			Header: RRHeader{
				RRType: 1,
				Class:  1,
				TTL:    86400,
			},
			Priority: 10,
			Weight:   5,
			Port:     80,
		}

		require.Equal(t, "dns_record", sr.Name)
		require.Equal(t, 1, sr.Header.RRType)
		require.Equal(t, 1, sr.Header.Class)
		require.Equal(t, 86400, sr.Header.TTL)
		require.Equal(t, 10, sr.Priority)
		require.Equal(t, 5, sr.Weight)
		require.Equal(t, 80, sr.Port)
		require.Equal(t, sr.RelationType(), model.SRVDNSRelation)
	})

	t.Run("Test successful JSON serialization of SRVDNSRelation", func(t *testing.T) {
		sr := SRVDNSRelation{
			Name: "dns_record",
			Header: RRHeader{
				RRType: 1,
				Class:  1,
				TTL:    86400,
			},
			Priority: 10,
			Weight:   5,
			Port:     80,
		}

		jsonData, err := sr.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"label":"dns_record", "header":{"rr_type":1, "class":1, "ttl":86400}, "priority":10, "weight":5, "port":80}`, string(jsonData))
	})
}
