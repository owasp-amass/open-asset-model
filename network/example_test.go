// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"fmt"
	"net/netip"
)

func ExampleIPAddress() {
	ip, _ := netip.ParseAddr("192.168.1.1")

	ipAdress := IPAddress{
		Address: ip,
		Type:    "IPv4",
	}

	fmt.Println(ipAdress.Address == ip)
	fmt.Println(ipAdress.Type == "IPv4")

	// Output:
	// true
	// true
}

func ExampleNetblock() {
	cidr, _ := netip.ParsePrefix("198.51.100.0/24")

	netblock := Netblock{
		Cidr: cidr,
		Type: "IPv4",
	}

	fmt.Println(netblock.Cidr == cidr)
	fmt.Println(netblock.Type == "IPv4")

	// Output:
	// true
	// true
}

func ExampleAutonomousSystem() {
	as := AutonomousSystem{
		Number: 64496,
	}

	fmt.Println(as.Number == 64496)

	// Output:
	// true
}
