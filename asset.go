package open_asset_model

type Asset interface {
	AssetType() AssetType
	JSON() ([]byte, error)
}

type AssetType string

const (
	IPAddress AssetType = "IPAddress"
	Netblock  AssetType = "Netblock"
	ASN       AssetType = "ASN"
	RIROrg    AssetType = "RIROrg"
	FQDN      AssetType = "FQDN"
)

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
