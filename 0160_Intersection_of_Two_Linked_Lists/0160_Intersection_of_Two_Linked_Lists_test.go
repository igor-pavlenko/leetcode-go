package leetcode_0160_Intersection_of_Two_Linked_Lists

import (
	"reflect"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {

	headA1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: nil,
					},
				},
			},
		},
	}
	headB1 := &ListNode{
		Val: 3,
		Next: headA1.Next,
	}

	headA2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: nil,
					},
				},
			},
		},
	}
	headB2 := &ListNode{
		Val: 3,
		Next: headA2.Next.Next.Next,
	}

	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				headA: headA1,
				headB: headB1,
			},
			want: headA1.Next,
		},
		{
			name: "Example 2",
			args: args{
				headA: headA2,
				headB: headB2,
			},
			want: headA2.Next.Next.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
			if got := getIntersectionNode2(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode2() = %v, want %v", got, tt.want)
			}
		})
	}

}
