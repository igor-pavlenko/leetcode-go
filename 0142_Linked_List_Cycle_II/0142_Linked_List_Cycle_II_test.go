package leetcode_0142_Linked_List_Cycle_II

import (
	"reflect"
	"testing"
)

func TestDetectCycle(t *testing.T) {

	head1 := &ListNode{
		Val: -3,
		Next: &ListNode{
			Val: -2,
			Next: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val:  0,
					Next: &ListNode{
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
			},
		},
	}
	head1.Next.Next.Next.Next.Next.Next.Next.Next = head1.Next.Next.Next

	head2 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
		},
	}
	head2.Next.Next.Next = head2.Next

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head: head1,
			},
			want: head1.Next.Next.Next,
		},
		{
			name: "Example 2",
			args: args{
				head: head2,
			},
			want: head2.Next,
		},
		{
			name: "Example 3",
			args: args{
				head: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
			if got := detectCycle2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
