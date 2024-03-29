package leetcode_0704_Binary_Search

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1", 
			args: args{
				nums:   []int{-1,0,3,5,9,12},
				target: 9,
			}, 
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				nums:   []int{-1,0,3,5,9,12},
				target: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
