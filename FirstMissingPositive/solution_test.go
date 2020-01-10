package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sorted array",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 6,
		},
		{
			name: "miss 3 in unsorted array",
			args: args{nums: []int{-2,8,-1,0,5,1,4,2,5,0}},
			want: 3,
		},
		{
			name: "miss 4 in unsorted array",
			args: args{nums: []int{1,2,3, 0}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstMissingPositive(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
