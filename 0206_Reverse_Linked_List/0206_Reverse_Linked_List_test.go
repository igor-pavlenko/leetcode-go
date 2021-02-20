package leetcode_0206_Reverse_Linked_List

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			want: ">4>3>2>1",
		},
		{
			name: "Example 2",
			args: args{
				head: nil,
			},
			want: "",
		},
		{
			name: "Example 3",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: ">1",
		},
		{
			name: "Example 4",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			want: ">3>2>1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.args.head)
			result := ""
			for got != nil {
				result = fmt.Sprintf("%s>%d", result, got.Val)
				got = got.Next
			}
			if tt.want != result {
				t.Errorf("removeElements() = %v, want %v", result, tt.want)
			}
		})
	}
}
