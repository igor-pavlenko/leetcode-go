package leetcode_0414_Third_Maximum_Number

import "testing"

func TestThirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{nums: []int{3, 2, 1}}, want: 1},
		{name: "Example 2", args: args{nums: []int{1, 2}}, want: 2},
		{name: "Example 3", args: args{nums: []int{2, 2, 2, 1, 3}}, want: 1},
		{name: "Example 4", args: args{nums: []int{1, 1, 2}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
