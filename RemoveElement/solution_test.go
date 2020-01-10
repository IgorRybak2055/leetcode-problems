package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple exemple",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "only val in slice",
			args:args{
				nums: []int{7, 7, 7, 7, 7},
				val:  7,
			},
			want: 0,
		},
		{
			name: "len(nums) = 1",
			args:args{
				nums: []int{7},
				val:  4,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			require.Equal(t, tt.want, got)
		})
	}
}
