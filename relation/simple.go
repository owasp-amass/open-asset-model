// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package relation

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// SimpleRelation represents a simple relation in the graph with no additional data required.
type SimpleRelation struct {
	Name string `json:"label"`
}

// Label implements the Relation interface.
func (r SimpleRelation) Label() string {
	return r.Name
}

// RelationType implements the Relation interface.
func (r SimpleRelation) RelationType() model.RelationType {
	return model.SimpleRelation
}

// JSON implements the Relation interface.
func (r SimpleRelation) JSON() ([]byte, error) {
	return json.Marshal(r)
}
