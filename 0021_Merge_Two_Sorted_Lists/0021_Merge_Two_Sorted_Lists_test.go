package leetcode_0021_Merge_Two_Sorted_Lists

import (
	"leetcode/lib"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				l1: &ListNode{
					Val:  5,
					Next: nil,
				},
				l2: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: "->1->2->4->5",
		},
		{
			name: "Example 2",
			args: args{
				l1: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: "->1->2->2->4->4->5",
		},
		{
			name: "Example 3",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := mergeTwoLists(tt.args.l1, tt.args.l2)
			got := lib.PrintList(head)
			if got != tt.want {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
