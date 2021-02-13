package leetcode_0283_Move_Zeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example 1", args: args{nums: []int{0, 1, 0, 3, 12}}, want: []int{1, 3, 12, 0, 0}},
		{name: "Example 2", args: args{nums: []int{2, 1, 0, 3, 0, 0}}, want: []int{2, 1, 3, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if moveZeroes(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
