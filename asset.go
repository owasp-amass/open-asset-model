package open_asset_model

type Asset interface {
	AssetType() AssetType
	JSON() ([]byte, error)
}

type AssetType string

const (
	IPAddress    AssetType = "IPAddress"
	Netblock     AssetType = "Netblock"
	ASN          AssetType = "ASN"
	RIROrg       AssetType = "RIROrg"
	FQDN         AssetType = "FQDN"
	WHOIS        AssetType = "WHOIS"
	Location     AssetType = "Location"
	Phone        AssetType = "Phone"
	Email        AssetType = "Email"
	Person       AssetType = "Person"
	Organization AssetType = "Organization"
)

var locationRels = map[string][]AssetType{}

var phoneRels = map[string][]AssetType{}

var emailRels = map[string][]AssetType{}

var whoisRels = map[string][]AssetType{
	"whois_registrar":  {Organization},
	"whois_registrant": {Organization},

	"whois_admin":          {Person},
	"whois_admin_phone":    {Phone},
	"whois_admin_email":    {Email},
	"whois_admin_location": {Location},

	"whois_tech":          {Person},
	"whois_tech_phone":    {Phone},
	"whois_tech_email":    {Email},
	"whois_tech_location": {Location},

	"whois_billing":          {Person},
	"whois_billing_phone":    {Phone},
	"whois_billing_email":    {Email},
	"whois_billing_location": {Location},

	// ?? do we add this as relations here or do we point to registrant
	// ?? then have registrant point to phone , email, location?
	"whois_registrant_phone":    {Phone},
	"whois_registrant_email":    {Email},
	"whois_registrant_location": {Location},

	"whois_nameserver": {FQDN},
	"whois_server":     {FQDN},

	// ?? does this need to be a relation? please check whois.go for more info on reseller.
	"whois_reseller": {Organization},
}

var personRels = map[string][]AssetType{
	"associated_with": {Organization},
	"phone_number":    {Phone},
	"email_address":   {Email},
	"location":        {Location},
}

var orgRels = map[string][]AssetType{
	"rir_org":      {RIROrg},
	"location":     {Location},
	"phone_number": {Phone},
	"email":        {Email},
}

var ipRels = map[string][]AssetType{}

var netblockRels = map[string][]AssetType{
	"contains": {IPAddress},
}

var asnRels = map[string][]AssetType{
	"announces":  {Netblock},
	"managed_by": {RIROrg},
}

var rirOrgRels = map[string][]AssetType{}

var fqdnRels = map[string][]AssetType{
	"a_record":     {IPAddress},
	"aaaa_record":  {IPAddress},
	"cname_record": {FQDN},
	"ns_record":    {FQDN},
	"ptr_record":   {FQDN},
	"mx_record":    {FQDN},
	"srv_record":   {FQDN, IPAddress},
	"node":         {FQDN},
}

// ValidRelationship returns true if the relation is valid in the taxonomy
// when outgoing from the source asset type to the destination asset type.
func ValidRelationship(source AssetType, relation string, destination AssetType) bool {
	var relations map[string][]AssetType

	switch source {
	case IPAddress:
		relations = ipRels
	case Netblock:
		relations = netblockRels
	case ASN:
		relations = asnRels
	case RIROrg:
		relations = rirOrgRels
	case FQDN:
		relations = fqdnRels
	case WHOIS:
		relations = whoisRels
	case Location:
		relations = locationRels
	case Phone:
		relations = phoneRels
	case Email:
		relations = emailRels
	case Person:
		relations = personRels
	case Organization:
		relations = orgRels
	default:
		return false
	}

	if atypes, ok := relations[relation]; ok {
		for _, atype := range atypes {
			if atype == destination {
				return true
			}
		}
	}
	return false
}
