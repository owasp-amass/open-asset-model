package contact_test

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/contact"
)

func TestPhone_AssetType(t *testing.T) {
	tests := []struct {
		name string
		p    Phone
		want model.AssetType
	}{
		{
			name: "phone asset type",
			p:    Phone{Type: "mobile", Phone: "123-456-7890"},
			want: model.Phone,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.AssetType(); got != tt.want {
				t.Errorf("Phone.AssetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhone_JSON(t *testing.T) {
	tests := []struct {
		name    string
		p       Phone
		want    []byte
		wantErr bool
	}{
		{
			name: "phone json",
			p:    Phone{Type: "mobile", Phone: "123-456-7890"},
			want: []byte(`{"type":"mobile","phone":"123-456-7890"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.JSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Phone.JSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Phone.JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
