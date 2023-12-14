package open_asset_model

type Asset interface {
	AssetType() AssetType
	JSON() ([]byte, error)
}

type AssetType string

const (
	IPAddress      AssetType = "IPAddress"
	Netblock       AssetType = "Netblock"
	ASN            AssetType = "ASN"
	RIROrg         AssetType = "RIROrg"
	FQDN           AssetType = "FQDN"
	WHOIS          AssetType = "WHOIS"
	Location       AssetType = "Location"
	Phone          AssetType = "Phone"
	Email          AssetType = "Email"
	Person         AssetType = "Person"
	Organization   AssetType = "Organization"
	Registrar      AssetType = "Registrar"
	Registrant     AssetType = "Registrant"
	TLSCertificate AssetType = "TLSCertificate"
)

var AssetList = []AssetType{
	IPAddress, Netblock, ASN, RIROrg, FQDN, WHOIS, Location,
	Phone, Email, Person, Organization, Registrar, Registrant,
}

var locationRels = map[string][]AssetType{}

var phoneRels = map[string][]AssetType{}

var emailRels = map[string][]AssetType{}

var whoisRels = map[string][]AssetType{
	"published_by": {Registrar},
	"name_server":  {FQDN},
	"reseller":     {Organization},

	"admin_org":      {Organization},
	"admin_person":   {Person},
	"admin_phone":    {Phone},
	"admin_email":    {Email},
	"admin_location": {Location},

	"tech_org":      {Organization},
	"tech_person":   {Person},
	"tech_phone":    {Phone},
	"tech_email":    {Email},
	"tech_location": {Location},

	"billing_org":      {Organization},
	"billing_person":   {Person},
	"billing_phone":    {Phone},
	"billing_email":    {Email},
	"billing_location": {Location},

	"registrant_org":      {Organization},
	"registrant_person":   {Person},
	"registrant_phone":    {Phone},
	"registrant_email":    {Email},
	"registrant_location": {Location},
}

var personRels = map[string][]AssetType{
	"phone_number":  {Phone},
	"email_address": {Email},
	"location":      {Location},
}

var orgRels = map[string][]AssetType{
	"rir_org":      {RIROrg},
	"location":     {Location},
	"phone_number": {Phone},
	"email":        {Email},
	"operates":     {Registrar},
}

var registrarRels = map[string][]AssetType{
	"abuse_email":  {Email},
	"abuse_phone":  {Phone},
	"whois_server": {FQDN},
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
	"registration": {WHOIS},
}

var tlscertRels = map[string][]AssetType{
	"common_name":               {FQDN},
	"subject_organization":      {Organization},
	"subject_organization_unit": {Organization},
	"subject_state_or_province": {Location},
	"subject_locality":          {Location},
	"subject_email_address":     {Email},
	"issuer":                    {FQDN},
	"issuer_organization":       {Organization},
	"issuer_organization_unit":  {Organization},
	"not_before":                {TLSCertificate},
	"not_after":                 {TLSCertificate},
	"subject_alt_names":         {FQDN},
	"signature_algorithm":       {TLSCertificate},
	"public_key_algorithm":      {TLSCertificate},
	"fingerprint_sha1":          {TLSCertificate},
	"fingerprint_sha256":        {TLSCertificate},
	"serial_number":             {TLSCertificate},
	"version":                   {TLSCertificate},
	"key_usage":                 {TLSCertificate},
	"extended_key_usage":        {TLSCertificate},
	"crl_distribution_points":   {TLSCertificate},
	"issuer_urls":               {FQDN},
	"ocsp_server":               {FQDN},
	"policies":                  {TLSCertificate},
	"subject_key_id":            {TLSCertificate},
	"authority_key_id":          {TLSCertificate},
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
	case Registrar:
		relations = registrarRels
	case TLSCertificate:
		relations = tlscertRels
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
