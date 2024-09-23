// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package registration

import (
	"net/netip"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
)

func TestIPNetRecordKey(t *testing.T) {
	want := "NET-150-154-0-0-1"
	as := IPNetRecord{Handle: want}

	if got := as.Key(); got != want {
		t.Errorf("IPNetRecord.Key() = %v, want %v", got, want)
	}
}

func TestIPNetRecordAssetType(t *testing.T) {
	var _ model.Asset = IPNetRecord{}       // Verify proper implementation of the Asset interface
	var _ model.Asset = (*IPNetRecord)(nil) // Verify the pointer properly implements the  Asset interface.

	w := IPNetRecord{}
	want := model.IPNetRecord

	if got := w.AssetType(); got != want {
		t.Errorf("IPNetRecord.AssetType() = %v, want %v", got, want)
	}
}
func TestIPNetRecord(t *testing.T) {
	record := IPNetRecord{
		CIDR:         netip.MustParsePrefix("150.154.0.0/16"),
		Handle:       "NET-150-154-0-0-1",
		StartAddress: netip.MustParseAddr("150.154.0.0"),
		EndAddress:   netip.MustParseAddr("150.154.255.255"),
		Type:         "IPv4",
		Name:         "REV-MVCC",
		Method:       "DIRECT ALLOCATION",
		ParentHandle: "NET-150-0-0-0-0",
		WhoisServer:  "whois.arin.net",
		CreatedDate:  "1991-05-20 04:00:00",
		UpdatedDate:  "2024-03-28 18:47:50",
		Status:       []string{"active"},
	}

	// Test AssetType method
	if record.AssetType() != model.IPNetRecord {
		t.Errorf("Expected asset type %s, but got %s", model.IPNetRecord, record.AssetType())
	}

	// Test JSON method
	expectedJSON := `{"cidr":"150.154.0.0/16","handle":"NET-150-154-0-0-1","start_address":"150.154.0.0","end_address":"150.154.255.255","type":"IPv4","name":"REV-MVCC","method":"DIRECT ALLOCATION","parent_handle":"NET-150-0-0-0-0","whois_server":"whois.arin.net","created_date":"1991-05-20 04:00:00","updated_date":"2024-03-28 18:47:50","status":["active"]}`
	json, err := record.JSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(json) != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, string(json))
	}
}
