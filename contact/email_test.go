package contact_test

import (
	"reflect"
	"testing"

	model "github.com/owasp-amass/open-asset-model"
	. "github.com/owasp-amass/open-asset-model/contact"
)

func TestEmail_AssetType(t *testing.T) {
	tests := []struct {
		name string
		e    Email
		want model.AssetType
	}{
		{
			name: "Test Email AssetType",
			e:    Email{Email: "test@test.com"},
			want: model.Email,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.AssetType(); got != tt.want {
				t.Errorf("Email.AssetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmail_JSON(t *testing.T) {
	tests := []struct {
		name    string
		e       Email
		want    []byte
		wantErr bool
	}{
		{
			name:    "Test Email JSON",
			e:       Email{Email: "test@test.com"},
			want:    []byte(`{"email":"test@test.com"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.JSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Email.JSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Email.JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
