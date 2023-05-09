package network_test

import (
	"fmt"
	"net/netip"

	. "github.com/owasp-amass/open-asset-model/network"
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

func ExampleRIROrganization() {
	riro := RIROrganization{
		Name:  "Google LLC",
		RIRId: "GOGL",
		RIR:   "ARIN",
	}

	fmt.Println(riro.Name == "Example")
	fmt.Println(riro.RIRId == "GOGL")
	fmt.Println(riro.RIR == "ARIN")

	// Output:
	// true
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
