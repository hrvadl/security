package ceasar

import "testing"

func TestCyclicShifterForward(t *testing.T) {
	t.Parallel()
	type fields struct {
		from  rune
		to    rune
		shift int
	}
	type args struct {
		target rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   rune
	}{
		{
			name: "Should shift correctly",
			fields: fields{
				from:  65,
				to:    90,
				shift: 4,
			},
			args: args{
				target: 65,
			},
			want: 69,
		},
		{
			name: "Should shift correctly when it's close to upper limit",
			fields: fields{
				from:  65,
				to:    90,
				shift: 4,
			},
			args: args{
				target: 88,
			},
			want: 66,
		},
		{
			name: "Should shift correctly when it's upper limit",
			fields: fields{
				from:  65,
				to:    90,
				shift: 4,
			},
			args: args{
				target: 90,
			},
			want: 68,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cs := &cyclicShifter{
				from:  tt.fields.from,
				to:    tt.fields.to,
				shift: tt.fields.shift,
			}
			if got := cs.Forward(tt.args.target); got != tt.want {
				t.Errorf("CyclicShifter.Forward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCyclicShifterBackward(t *testing.T) {
	t.Parallel()
	type fields struct {
		from  rune
		to    rune
		shift int
	}
	type args struct {
		target rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   rune
	}{
		{
			name: "Should shift correctly",
			fields: fields{
				from:  65,
				to:    90,
				shift: 4,
			},
			args: args{
				target: 75,
			},
			want: 71,
		},
		{
			name: "Should shift correctly when it's close to the lower limit",
			fields: fields{
				from:  65,
				to:    90,
				shift: 4,
			},
			args: args{
				target: 67,
			},
			want: 89,
		},
		{
			name: "Should shift correctly when it's lower limit",
			fields: fields{
				from:  65,
				to:    90,
				shift: 4,
			},
			args: args{
				target: 65,
			},
			want: 87,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cs := &cyclicShifter{
				from:  tt.fields.from,
				to:    tt.fields.to,
				shift: tt.fields.shift,
			}
			if got := cs.Backward(tt.args.target); got != tt.want {
				t.Errorf("CyclicShifter.Backward() = %v, want %v", got, tt.want)
			}
		})
	}
}
