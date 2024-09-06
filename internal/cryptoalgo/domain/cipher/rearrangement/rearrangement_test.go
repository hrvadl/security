package rearrangement

import (
	"reflect"
	"testing"
)

func TestWithRearrangementEncrypt(t *testing.T) {
	t.Parallel()
	type fields struct {
		key []int
	}
	type args struct {
		msg []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Should encrypt with rearrangement correctly",
			fields: fields{
				key: []int{4, 3, 2, 1},
			},
			args: args{
				msg: []byte("text"),
			},
			want: []byte("txet"),
		},
		{
			name: "Should throw error if key's length does not match text's length",
			fields: fields{
				key: []int{4, 3, 2, 1},
			},
			args: args{
				msg: []byte("hello world!"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			wr := &Cipher{
				key: tt.fields.key,
			}

			got, err := wr.Encrypt(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithRearrangement.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRearrangement.Encrypt() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestWithRearrangementDecrypt(t *testing.T) {
	t.Parallel()
	type fields struct {
		key []int
	}
	type args struct {
		msg []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Should encrypt with rearrangement correctly",
			fields: fields{
				key: []int{4, 3, 2, 1},
			},
			args: args{
				msg: []byte("txet"),
			},
			want: []byte("text"),
		},
		{
			name: "Should throw error if key's length does not match text's length",
			fields: fields{
				key: []int{4, 3, 2, 1},
			},
			args: args{
				msg: []byte("hello world!"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			wr := &Cipher{
				key: tt.fields.key,
			}

			got, err := wr.Decrypt(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithRearrangement.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRearrangement.Encrypt() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
