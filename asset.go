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
	Address      AssetType = "Address"
	Phone        AssetType = "Phone"
	Email        AssetType = "Email"
	Name         AssetType = "Name"
	Role         AssetType = "Role"
	Organization AssetType = "Organization"
)

var addressRels = map[string][]AssetType{
	"physical_address": {Organization, Name, Role},
}

var phoneRels = map[string][]AssetType{
	"phone_number": {Organization, Name, Role},
}

var emailRels = map[string][]AssetType{
	"email_address": {Organization, Name, Role, FQDN},
}

var whoisRels = map[string][]AssetType{
	"registrar":    {Organization},
	"obtained_by":  {Email, Phone, Address},
	"whois_record": {FQDN, Netblock},
}

var nameRels = map[string][]AssetType{
	"associated_with": {Organization},
	"role":            {Role},
}

var orgRels = map[string][]AssetType{
	"rir_name": {RIROrg},
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
	case Address:
		relations = addressRels
	case Phone:
		relations = phoneRels
	case Email:
		relations = emailRels
	case Name:
		relations = nameRels
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
