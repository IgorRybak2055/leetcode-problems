package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "with target=0",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name: "with target=5",
			args: args{
				nums:   []int{7, 6, 0, -1, 0, -2, 2},
				target: 5,
			},
			want: [][]int{{-2, -1, 2, 6}, {-2, 0, 0, 7}, {-1, 0, 0, 6}},
		},
		{
			name: "with nil answer",
			args: args{
				nums:   []int{7, 6, 0, -1, 0, -2, 2},
				target: 2,
			},
			want: nil,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.args.nums, tt.args.target)
			require.EqualValues(t, tt.want, got)
		})
	}
}
