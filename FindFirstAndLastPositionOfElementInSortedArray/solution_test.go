package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "correct dataset 1",
			args: args{
				nums:   []int{5, 7, 8, 8, 8, 10, 10},
				target: 8,
			},
			want: []int{2, 4},
		},
		{
			name: "correct dataset 2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
		{
			name: "correct dataset 3",
			args: args{
				nums:   []int{1, 1, 1, 3, 4, 5, 7, 8, 8, 8},
				target: 8,
			},
			want: []int{7, 9},
		},
		{
			name: "correct dataset 4",
			args: args{
				nums:   []int{1, 1, 1, 3, 4, 5, 7, 8, 9, 9, 10, 11},
				target: 8,
			},
			want: []int{7, 7},
		},
		{
			name: "incorrect dataset",
			args: args{
				nums:   []int{1, 1, 1, 3, 4, 5, 7, 10, 15, 18, 19, 21, 21},
				target: 8,
			},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRange(tt.args.nums, tt.args.target)
			require.Equal(t, tt.want, got)
		})
	}
}
