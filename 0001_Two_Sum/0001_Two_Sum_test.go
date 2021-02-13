package leetcode_0001_Two_Sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example 1", args: args{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{name: "Example 2", args: args{nums: []int{3, 2, 4}, target: 6}, want: []int{1, 2}},
		{name: "Example 3", args: args{nums: []int{3, 3}, target: 6}, want: []int{0, 1}},
		{name: "Example 100", args: args{nums: []int{}, target: 0}, want: nil},
		{name: "Example 101", args: args{nums: []int{1}, target: 0}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("0001_Two_Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
