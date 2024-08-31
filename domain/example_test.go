// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package domain

import (
	"fmt"
)

func ExampleFQDN() {
	fqdn := FQDN{
		Name: "foo.example.com",
	}

	fmt.Println(fqdn.Name == "foo.example.com")

	// Output:
	// true
}
