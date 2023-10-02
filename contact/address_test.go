package contact_test

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/contact"
)

func TestAddress_AssetType(t *testing.T) {
	tests := []struct {
		name string
		a    Address
		want model.AssetType
	}{
		{
			name: "returns Address asset type",
			a: Address{
				Street: "123 Main St",
			},
			want: model.Address,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AssetType(); got != tt.want {
				t.Errorf("Address.AssetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_JSON(t *testing.T) {
	tests := []struct {
		name    string
		a       Address
		want    []byte
		wantErr bool
	}{
		{
			name: "returns JSON representation of address",
			a: Address{
				Street:        "123 Main St",
				City:          "Anytown",
				StateProvince: "CA",
				PostalCode:    "12345",
				Country:       "USA",
			},
			want:    []byte(`{"street":"123 Main St","city":"Anytown","state/province":"CA","postal_code":"12345","country":"USA"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.JSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Address.JSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
