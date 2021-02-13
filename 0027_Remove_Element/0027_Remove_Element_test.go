package leetcode_0027_Remove_Element

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantSlice []int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want:      2,
			wantSlice: []int{2, 2},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want:      5,
			wantSlice: []int{0, 1, 3, 0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val);
				got != tt.want || !reflect.DeepEqual(tt.args.nums[0:got], tt.wantSlice) {
				t.Errorf("removeElement() = %v, want %v, nums = %v, want %v", got, tt.want, tt.args.nums[0:got], tt.wantSlice)
			}
		})
	}
}
