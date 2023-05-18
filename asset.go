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
