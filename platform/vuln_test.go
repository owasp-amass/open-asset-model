// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package platform

import (
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	"github.com/stretchr/testify/require"
)

func TestVulnPropertyName(t *testing.T) {
	want := "CVE-2019-00001"
	sp := VulnProperty{
		ID:          "CVE-2019-00001",
		Description: "In macOS before 2.12.6, there is a vulnerability in the RPC...",
		Source:      "Tenable",
		Category:    "Firewall",
		Enumeration: "CVE",
		Reference:   "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-6111",
	}

	if got := sp.Name(); got != want {
		t.Errorf("VulnProperty.Name() = %v, want %v", got, want)
	}
}

func TestVulnPropertyValue(t *testing.T) {
	want := "In macOS before 2.12.6, there is a vulnerability in the RPC..."
	sp := VulnProperty{
		ID:          "CVE-2019-00001",
		Description: "In macOS before 2.12.6, there is a vulnerability in the RPC...",
		Source:      "Tenable",
		Category:    "Firewall",
		Enumeration: "CVE",
		Reference:   "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-6111",
	}

	if got := sp.Value(); got != want {
		t.Errorf("VulnProperty.Value() = %v, want %v", got, want)
	}
}

func TestVulnPropertyImplementsProperty(t *testing.T) {
	var _ model.Property = VulnProperty{}       // Verify proper implementation of the Property interface
	var _ model.Property = (*VulnProperty)(nil) // Verify *VulnProperty properly implements the Property interface.
}

func TestVulnProperty(t *testing.T) {
	t.Run("Test successful creation of VulnProperty", func(t *testing.T) {
		sp := VulnProperty{
			ID:          "CVE-2019-00001",
			Description: "In macOS before 2.12.6, there is a vulnerability in the RPC...",
			Source:      "Tenable",
			Category:    "Firewall",
			Enumeration: "CVE",
			Reference:   "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-6111",
		}

		require.Equal(t, "CVE-2019-00001", sp.ID)
		require.Equal(t, "In macOS before 2.12.6, there is a vulnerability in the RPC...", sp.Description)
		require.Equal(t, "Tenable", sp.Source)
		require.Equal(t, "Firewall", sp.Category)
		require.Equal(t, "CVE", sp.Enumeration)
		require.Equal(t, "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-6111", sp.Reference)
		require.Equal(t, sp.PropertyType(), model.VulnProperty)
	})

	t.Run("Test successful JSON serialization of VulnProperty", func(t *testing.T) {
		sp := VulnProperty{
			ID:          "CVE-2019-00001",
			Description: "foobar",
			Source:      "Tenable",
			Category:    "Firewall",
			Enumeration: "CVE",
			Reference:   "URL",
		}

		jsonData, err := sp.JSON()

		require.NoError(t, err)
		require.JSONEq(t, `{"id":"CVE-2019-00001", "desc":"foobar", "source":"Tenable", "category":"Firewall", "enum":"CVE", "ref":"URL"}`, string(jsonData))
	})
}
