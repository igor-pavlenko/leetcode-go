package leetcode_0007_Reverse_Integer

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{x: 1230}, want: 321},
		{name: "Example 2", args: args{x: -123}, want: -321},
		{name: "Example 3", args: args{x: 120}, want: 21},
		{name: "Example 4", args: args{x: 0}, want: 0},
		{name: "Example 5", args: args{x: 1534236469}, want: 0},
		{name: "Example 100", args: args{x: math.MaxInt32 + 1}, want: 0},
		{name: "Example 101", args: args{x: math.MinInt32 - 1}, want: 0},
		{name: "Example 102", args: args{x: -2147483412}, want: -2143847412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
