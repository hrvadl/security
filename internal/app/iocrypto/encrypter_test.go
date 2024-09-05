package iocrypto

import (
	"bytes"
	"io"
	"testing"

	"github.com/hrvadl/security/internal/domain/cipher/rearrangement"
)

func TestEncrypterEncrypt(t *testing.T) {
	type fields struct {
		in     io.Reader
		out    *bytes.Buffer
		cipher CipherEncryptSource
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Should encrypt correctly",
			fields: fields{
				in:     bytes.NewBufferString("hello world!!!!"),
				out:    bytes.NewBufferString(""),
				cipher: rearrangement.NewCipher([]int{5, 4, 3, 2, 1}),
			},
			want: []byte("ollehlrow !!!!d"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encrypter{
				in:     tt.fields.in,
				out:    tt.fields.out,
				cipher: tt.fields.cipher,
			}

			if err := e.Encrypt(); (err != nil) != tt.wantErr {
				t.Errorf("Encrypter.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !bytes.Equal(tt.want, tt.fields.out.Bytes()) {
				t.Errorf(
					"Encrypter.Encrypt() want = %v, got %v",
					string(tt.want),
					tt.fields.out.String(),
				)
			}
		})
	}
}
