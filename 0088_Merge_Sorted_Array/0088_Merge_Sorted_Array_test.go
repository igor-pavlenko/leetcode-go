package leetcode_0088_Merge_Sorted_Array

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example 1", args: args{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
		}, want: []int{1, 2, 2, 3, 5, 6}},
		{name: "Example 2", args: args{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
		}, want: []int{1}},
		{name: "Example 3", args: args{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
		}, want: []int{1}},
		{name: "Example 4", args: args{
			nums1: []int{4, 5, 6, 0, 0, 0},
			m:     3,
			nums2: []int{1, 2, 3},
			n:     3,
		}, want: []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
