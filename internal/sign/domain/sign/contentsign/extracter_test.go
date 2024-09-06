package contentsign

import (
	"reflect"
	"testing"
)

func TestExtracterExtractSign(t *testing.T) {
	t.Parallel()
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   []byte
		wantErr bool
	}{
		{
			name: "Should extract signature correctly",
			args: args{
				data: []byte("some data\n---SIGNATURE---\nthis is a signature"),
			},
			want:    []byte("some data"),
			want1:   []byte("this is a signature"),
			wantErr: false,
		},
		{
			name: "Should extract signature correctly",
			args: args{
				data: []byte("some data\nand even more\n---SIGNATURE---\nthis is a signature"),
			},
			want:    []byte("some data\nand even more"),
			want1:   []byte("this is a signature"),
			wantErr: false,
		},
		{
			name: "Should return error if no signature header is present",
			args: args{
				data: []byte("some data\n"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := &Extracter{}
			got, got1, err := e.ExtractSign(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Extracter.ExtractSign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extracter.ExtractSign() got = %v, want %v", string(got), string(tt.want))
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf(
					"Extracter.ExtractSign() got1 = %v, want %v",
					string(got1),
					string(tt.want1),
				)
			}
		})
	}
}
