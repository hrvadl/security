package contentsign

import (
	"reflect"
	"testing"
)

func TestAppenderAppendSign(t *testing.T) {
	t.Parallel()
	type args struct {
		data []byte
		sign []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Should append signature correctly",
			args: args{
				data: []byte("some data"),
				sign: []byte("some signature"),
			},
			want: []byte("some data\n---SIGNATURE---\nsome signature"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := &Appender{}
			if got := a.AppendSign(tt.args.data, tt.args.sign); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Appender.AppendSign() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
