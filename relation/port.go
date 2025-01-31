// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package relation

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// PortRelation is a relation in the graph representing an open port.
type PortRelation struct {
	Name       string `json:"label"`
	PortNumber int    `json:"port_number"`
	Protocol   string `json:"protocol"`
}

// Label implements the Relation interface.
func (r PortRelation) Label() string {
	return r.Name
}

// RelationType implements the Relation interface.
func (r PortRelation) RelationType() model.RelationType {
	return model.PortRelation
}

// JSON implements the Relation interface.
func (r PortRelation) JSON() ([]byte, error) {
	return json.Marshal(r)
}
