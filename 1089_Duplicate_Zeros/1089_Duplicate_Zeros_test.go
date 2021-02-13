package leetcode_1089_Duplicate_Zeros

import (
	"reflect"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example 1", args: args{arr: []int{1, 0, 2, 3, 0, 4, 5, 0}}, want: []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{name: "Example 1", args: args{arr: []int{0, 1, 2, 3}}, want: []int{0, 0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := duplicateZeros(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("duplicateZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
