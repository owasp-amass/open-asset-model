// Copyright © by Jeff Foley 2017-2026. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package platform

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// StorageBucket represents an object-storage bucket — an Amazon S3, Google
// Cloud Storage or Azure Blob container addressed by name within a provider's
// global namespace, rather than a service listening on a host port. The same
// bucket may be reached through several host names (virtual-hosted, website and
// CNAME endpoints); its identity is the provider and bucket name, not any one
// endpoint. A bucket that does not resolve at the provider (Exists is false) is
// a claimable subdomain-takeover target.
//
// It should support relationships such as the following:
// - Provider (the cloud Organization that operates the namespace)
// - Bucket host (the FQDN reaching it, via the FQDN "bucket" relation)
type StorageBucket struct {
	ID       string `json:"unique_id"`         // "<service>:<bucket>", e.g. "s3:example-public"
	Name     string `json:"bucket_name"`       // the bucket name within the provider namespace
	Provider string `json:"provider"`          // e.g. "Amazon Web Services", "Google Cloud", "Microsoft Azure"
	Service  string `json:"service,omitempty"` // "s3" | "gcs" | "azure-blob"
	Exists   bool   `json:"exists"`            // resolves at the provider; false means claimable
	Public   bool   `json:"public"`            // world-listable
}

// Key implements the Asset interface.
func (b StorageBucket) Key() string {
	return b.ID
}

// AssetType implements the Asset interface.
func (b StorageBucket) AssetType() model.AssetType {
	return model.StorageBucket
}

// JSON implements the Asset interface.
func (b StorageBucket) JSON() ([]byte, error) {
	return json.Marshal(b)
}
