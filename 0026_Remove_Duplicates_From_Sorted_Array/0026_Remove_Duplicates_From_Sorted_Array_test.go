package _026_Remove_Duplicates_From_Sorted_Array

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example 1", args: args{nums: []int{1, 1, 2}}, want: []int{1, 2}},
		{name: "Example 2", args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, want: []int{0, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example 1", args: args{nums: []int{1, 1, 2}}, want: []int{1, 2}},
		{name: "Example 1", args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, want: []int{0, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}
