package domain_test

import (
	"fmt"

	. "github.com/owasp-amass/open-asset-model/domain"
)

func ExampleFQDN() {
	fqdn := FQDN{
		Name: "foo.example.com",
	}

	fmt.Println(fqdn.Name == "foo.example.com")

	// Output:
	// true
}
