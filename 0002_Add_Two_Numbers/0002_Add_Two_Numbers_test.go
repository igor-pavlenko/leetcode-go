package leestcode_0002_Add_Two_Numbers

import (
	"fmt"
	"testing"
)

func printList(head *ListNode) string {
	result := ""
	for head != nil {
		result = fmt.Sprintf("%s->%d", result, head.Val)
		head = head.Next
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
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
					Val:  9,
					Next: &ListNode{
						Val:  9,
						Next: &ListNode{
							Val:  9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
				l2: &ListNode{
					Val:  9,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
			want: "->8->9->0->0->1",
		},
		{
			name: "Example 2",
			args: args{
				l1: &ListNode{
					Val:  0,
					Next: nil,
				},
				l2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: "->0",
		},
		{
			name: "Example 3",
			args: args{
				l1: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val:  5,
					Next: &ListNode{
						Val:  6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: "->7->0->8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2);
			gotStr := printList(got)
			if gotStr != tt.want {
				t.Errorf("addTwoNumbers() = %s, want %s", gotStr, tt.want)
			}
		})
	}
}
