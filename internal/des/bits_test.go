package des

import (
	"reflect"
	"testing"
)

func Test_toBits(t *testing.T) {
	t.Parallel()
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    []Bit
		wantErr bool
	}{
		{
			name: "Should convert string to bits correctly",
			args: args{
				text: "hellostr",
			},
			want: []Bit{
				0, 1, 1, 0, 1, 0, 0, 0,
				0, 1, 1, 0, 0, 1, 0, 1,
				0, 1, 1, 0, 1, 1, 0, 0,
				0, 1, 1, 0, 1, 1, 0, 0,
				0, 1, 1, 0, 1, 1, 1, 1,
				0, 1, 1, 1, 0, 0, 1, 1,
				0, 1, 1, 1, 0, 1, 0, 0,
				0, 1, 1, 1, 0, 0, 1, 0,
			},
		},
		{
			name: "Should convert string to bits correctly",
			args: args{
				text: "thisis8b",
			},
			want: []Bit{
				0, 1, 1, 1, 0, 1, 0, 0,
				0, 1, 1, 0, 1, 0, 0, 0,
				0, 1, 1, 0, 1, 0, 0, 1,
				0, 1, 1, 1, 0, 0, 1, 1,
				0, 1, 1, 0, 1, 0, 0, 1,
				0, 1, 1, 1, 0, 0, 1, 1,
				0, 0, 1, 1, 1, 0, 0, 0,
				0, 1, 1, 0, 0, 0, 1, 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := toBits(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("toBits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fromBits(t *testing.T) {
	t.Parallel()
	type args struct {
		b []Bit
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should convert from bits correctly",
			args: args{
				b: []Bit{
					0, 1, 1, 1, 0, 1, 0, 0,
					0, 1, 1, 0, 1, 0, 0, 0,
					0, 1, 1, 0, 1, 0, 0, 1,
					0, 1, 1, 1, 0, 0, 1, 1,
					0, 1, 1, 0, 1, 0, 0, 1,
					0, 1, 1, 1, 0, 0, 1, 1,
					0, 0, 1, 1, 1, 0, 0, 0,
					0, 1, 1, 0, 0, 0, 1, 0,
				},
			},
			want: "thisis8b",
		},
		{
			name: "Should convert string to bits correctly",
			args: args{
				b: []Bit{
					0, 1, 1, 0, 1, 0, 0, 0,
					0, 1, 1, 0, 0, 1, 0, 1,
					0, 1, 1, 0, 1, 1, 0, 0,
					0, 1, 1, 0, 1, 1, 0, 0,
					0, 1, 1, 0, 1, 1, 1, 1,
					0, 1, 1, 1, 0, 0, 1, 1,
					0, 1, 1, 1, 0, 1, 0, 0,
					0, 1, 1, 1, 0, 0, 1, 0,
				},
			},
			want: "hellostr",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := fromBits(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("fromBits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fromBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
