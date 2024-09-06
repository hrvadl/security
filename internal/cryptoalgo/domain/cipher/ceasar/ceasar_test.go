package ceasar

import (
	"reflect"
	"testing"
)

func TestCipherEncrypt(t *testing.T) {
	t.Parallel()
	type fields struct {
		shifter Shifter
	}
	type args struct {
		msg []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "Should encrypt correctly",
			fields: fields{
				shifter: NewShiftStrategy(4),
			},
			args: args{
				msg: []byte("hello world"),
			},
			want: []byte("lipps asvph"),
		},
		{
			name: "Should encrypt correctly",
			fields: fields{
				shifter: NewShiftStrategy(4),
			},
			args: args{
				msg: []byte("how u doin"),
			},
			want: []byte("lsa y hsmr"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &Cipher{
				shifter: tt.fields.shifter,
			}
			if got := c.Encrypt(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cipher.Encrypt() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestCipherDecrypt(t *testing.T) {
	t.Parallel()
	type fields struct {
		shifter Shifter
	}
	type args struct {
		msg []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "Should encrypt correctly",
			fields: fields{
				shifter: NewShiftStrategy(4),
			},
			args: args{
				msg: []byte("lipps asvph"),
			},
			want: []byte("hello world"),
		},
		{
			name: "Should encrypt correctly",
			fields: fields{
				shifter: NewShiftStrategy(4),
			},
			args: args{
				msg: []byte("lsa y hsmr"),
			},
			want: []byte("how u doin"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &Cipher{
				shifter: tt.fields.shifter,
			}
			if got := c.Decrypt(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cipher.Decrypt() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
