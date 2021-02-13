package leetcode_0941_Valid_Mountain_Array

import "testing"

func TestValidMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example 1", args: args{arr: []int{2, 1}}, want: false},
		{name: "Example 2", args: args{arr: []int{3, 5, 5}}, want: false},
		{name: "Example 3", args: args{arr: []int{0, 3, 2, 1}}, want: true},
		{name: "Example 4", args: args{arr: []int{0, 2, 3, 3, 5, 2, 1, 0}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
