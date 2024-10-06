package des

import (
	"reflect"
	"testing"
)

func TestStringToBits(t *testing.T) {
	t.Parallel()
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []Bit
	}{
		{
			name: "Should convert string to bits correctly",
			args: args{
				text: "hellostr",
			},
			want: []Bit{
				"0", "1", "1", "0", "1", "0", "0", "0",
				"0", "1", "1", "0", "0", "1", "0", "1",
				"0", "1", "1", "0", "1", "1", "0", "0",
				"0", "1", "1", "0", "1", "1", "0", "0",
				"0", "1", "1", "0", "1", "1", "1", "1",
				"0", "1", "1", "1", "0", "0", "1", "1",
				"0", "1", "1", "1", "0", "1", "0", "0",
				"0", "1", "1", "1", "0", "0", "1", "0",
			},
		},
		{
			name: "Should convert string to bits correctly",
			args: args{
				text: "thisis8b",
			},
			want: []Bit{
				"0", "1", "1", "1", "0", "1", "0", "0",
				"0", "1", "1", "0", "1", "0", "0", "0",
				"0", "1", "1", "0", "1", "0", "0", "1",
				"0", "1", "1", "1", "0", "0", "1", "1",
				"0", "1", "1", "0", "1", "0", "0", "1",
				"0", "1", "1", "1", "0", "0", "1", "1",
				"0", "0", "1", "1", "1", "0", "0", "0",
				"0", "1", "1", "0", "0", "0", "1", "0",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := stringToBits(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
