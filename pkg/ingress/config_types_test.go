package ingress

import "testing"

func TestAuthIDPConfigOIDC_getInfoFromDiscoveryUrl(t *testing.T) {
	type fields struct {
		DiscoveryEndpoint                string
		Issuer                           string
		AuthorizationEndpoint            string
		TokenEndpoint                    string
		UserInfoEndpoint                 string
		SecretName                       string
		AuthenticationRequestExtraParams map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "with discovery url",
			fields: fields{
				DiscoveryEndpoint: "https://example.com/.well-known/openid-configuration",
				AuthenticationRequestExtraParams: map[string]string{
					"key": "value",
				},
			},
			wantErr: false,
		},
		{
			name: "without discovery url",
			fields: fields{
				Issuer:                "https://example.com",
				AuthorizationEndpoint: "https://authorization.example.com",
				TokenEndpoint:         "https://token.example.com",
				UserInfoEndpoint:      "https://userinfo.example.com",
				SecretName:            "my-k8s-secret",
				AuthenticationRequestExtraParams: map[string]string{
					"key": "value",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthIDPConfigOIDC{
				DiscoveryEndpoint:                tt.fields.DiscoveryEndpoint,
				Issuer:                           tt.fields.Issuer,
				AuthorizationEndpoint:            tt.fields.AuthorizationEndpoint,
				TokenEndpoint:                    tt.fields.TokenEndpoint,
				UserInfoEndpoint:                 tt.fields.UserInfoEndpoint,
				SecretName:                       tt.fields.SecretName,
				AuthenticationRequestExtraParams: tt.fields.AuthenticationRequestExtraParams,
			}
			if err := a.getInfoFromDiscoveryUrl(); (err != nil) != tt.wantErr {
				t.Errorf("AuthIDPConfigOIDC.getInfoFromDiscoveryUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
