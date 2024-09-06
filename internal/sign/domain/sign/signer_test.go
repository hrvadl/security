package sign

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/hrvadl/security/internal/sign/domain/hash"
)

func TestSignerSign(t *testing.T) {
	t.Parallel()
	type fields struct {
		key    *rsa.PrivateKey
		hasher Hasher
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should create signature correctly",
			fields: fields{
				key:    mustGenerateKey(t),
				hasher: hash.NewHasher(),
			},
			args: args{
				data: []byte("test"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Signer{
				key:    tt.fields.key,
				hasher: tt.fields.hasher,
			}
			_, err := s.SignToBase64(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Signer.Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func mustGenerateKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Fatalf("Failed to generate rsa key: %v", err)
	}

	return key
}
