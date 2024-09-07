package analysis

import "testing"

func TestCeasarKeyDecryptorGetKeyFor(t *testing.T) {
	t.Parallel()
	type args struct {
		source    []byte
		encrypted []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Should get key correctly",
			args: args{
				source:    []byte("hello world"),
				encrypted: []byte("lipps asvph"),
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "Should get key correctly",
			args: args{
				source:    []byte("hello world"),
				encrypted: []byte("mjqqt btwqi"),
			},
			want:    5,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &CaesarKeyDecryptor{}
			got, err := c.GetKeyFor(tt.args.source, tt.args.encrypted)
			if (err != nil) != tt.wantErr {
				t.Errorf("CeasarKeyDecryptor.GetKeyFor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CeasarKeyDecryptor.GetKeyFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
