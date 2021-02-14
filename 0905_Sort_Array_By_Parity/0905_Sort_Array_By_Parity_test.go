package leetcode_0905_Sort_Array_By_Parity

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example 1", args: args{A: []int{3, 1, 2, 4}}, want: []int{2, 4, 3, 1}},
		{name: "Example 2", args: args{A: []int{2, 4, 6, 0, 1, 3, 5}}, want: []int{2, 4, 6, 0, 1, 3, 5}},
		{name: "Example 3", args: args{A: []int{1, 3, 5}}, want: []int{1, 3, 5}},
		{name: "Example 4", args: args{A: []int{2, 2}}, want: []int{2, 2}},
		{name: "Example 5", args: args{A: []int{1}}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
