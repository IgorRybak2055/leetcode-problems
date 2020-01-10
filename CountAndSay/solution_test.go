package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "with value = 5" ,
			args: args{n:5},
			want: "111221",
		},
		{
			name: "with value = 1" ,
			args: args{n:1},
			want: "1",
		},
		{
			name: "with value = 3" ,
			args: args{n:3},
			want: "21",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countAndSay(tt.args.n)

			require.Equal(t, tt.want, got)
		})
	}
}