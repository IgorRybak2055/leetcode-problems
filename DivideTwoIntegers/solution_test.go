package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "MaxInt",
			args: args{
				dividend: -2147483648,
				divisor:  -1,
			},
			want: math.MaxInt32,
		},
		{
			name: "MinInt",
			args:args{
				dividend: 2147483649,
				divisor:  -1,
			},
			want: math.MinInt32,
		},
		{
			name: "check division by zero",
			args:args{
				dividend: 7,
				divisor:  0,
			},
			want: 0,
		},
		{
			name: "check usual division",
			args:args{
				dividend: 7,
				divisor:  0,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divide(tt.args.dividend, tt.args.divisor)
			assert.Equal(t, tt.want, got)
		})
	}
}
