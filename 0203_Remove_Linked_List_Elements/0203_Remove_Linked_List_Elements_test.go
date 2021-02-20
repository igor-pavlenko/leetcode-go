package leetcode_0203_Remove_Linked_List_Elements

import (
	"fmt"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
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
				val:  2,
			},
			want: ">1>3>4",
		},
		{
			name: "Example 2",
			args: args{
				head: nil,
				val:  1,
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
				val:  1,
			},
			want: "",
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
				val:  1,
			},
			want: ">2>3",
		},
		{
			name: "Example 5",
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
				val:  3,
			},
			want: ">1>2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElements(tt.args.head, tt.args.val)
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
