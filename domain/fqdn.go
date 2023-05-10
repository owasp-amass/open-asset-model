package domain

// FQDN represents a Fully Qualified Domain Name.
type FQDN struct {
	// Name is the domain name part of the FQDN.
	// It should be a valid domain name, such as "example" or "subdomain.example.com".
	Name string `json:"name"`
}
