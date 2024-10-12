package main

import (
	"reflect"
	"testing"
)

// I've used https://devtoolcafe.com/tools/des tool
// to verify DES encryption results.
func Test_encrypt(t *testing.T) {
	t.Parallel()
	type args struct {
		key  []byte
		text []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("hellokey"),
				text: []byte("hellostr"),
			},
			want: []byte("1GCLngrmrUI="),
		},
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("hellostr"),
				text: []byte("hellostr"),
			},
			want: []byte("k+Q66w/zdLk="),
		},
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("thatskey"),
				text: []byte("thistext"),
			},
			want: []byte("q2jKkYCtIiY="),
		},
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("thiskey!"),
				text: []byte("thistext"),
			},
			want: []byte("3Gp6Q3S99KQ="),
		},
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("12345678"),
				text: []byte("thistext"),
			},
			want: []byte("rPGTyCJo1gg="),
		},
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("12345678"),
				text: []byte("12345678"),
			},
			want: []byte("ltACiHjVjIk="),
		},
		{
			name: "Should return error if key is less than 8 bytes",
			args: args{
				key:  []byte{1},
				text: make([]byte, allowedTextLength),
			},
			wantErr: true,
		},
		{
			name: "Should return error if key is more than 8 bytes",
			args: args{
				key:  make([]byte, 666),
				text: make([]byte, allowedTextLength),
			},
			wantErr: true,
		},
		{
			name: "Should return error if text is less than 8 bytes",
			args: args{
				key:  make([]byte, allowedTextLength),
				text: []byte{1},
			},
			wantErr: true,
		},
		{
			name: "Should return error if text is more than 8 bytes",
			args: args{
				key:  make([]byte, allowedTextLength),
				text: make([]byte, 666),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := encrypt(tt.args.key, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encrypt() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func Benchmark_encrypt(b *testing.B) {
	b.ReportAllocs()
	type args struct {
		key  []byte
		text []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("hellokey"),
				text: []byte("hellostr"),
			},
			want: []byte("1GCLngrmrUI="),
		},
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("hellostr"),
				text: []byte("hellostr"),
			},
			want: []byte("k+Q66w/zdLk="),
		},
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("thatskey"),
				text: []byte("thistext"),
			},
			want: []byte("q2jKkYCtIiY="),
		},
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("thiskey!"),
				text: []byte("thistext"),
			},
			want: []byte("3Gp6Q3S99KQ="),
		},
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("12345678"),
				text: []byte("thistext"),
			},
			want: []byte("rPGTyCJo1gg="),
		},
		{
			name: "Should encrypt text correctly",
			args: args{
				key:  []byte("12345678"),
				text: []byte("12345678"),
			},
			want: []byte("ltACiHjVjIk="),
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			got, err := encrypt(tt.args.key, tt.args.text)
			if (err != nil) != tt.wantErr {
				b.Errorf("encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				b.Errorf("encrypt() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
