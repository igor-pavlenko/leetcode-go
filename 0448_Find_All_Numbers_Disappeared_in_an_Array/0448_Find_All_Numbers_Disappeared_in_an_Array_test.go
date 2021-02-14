package leetcode_0448_Find_All_Numbers_Disappeared_in_an_Array

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example 1", args: args{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}}, want: []int{5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
