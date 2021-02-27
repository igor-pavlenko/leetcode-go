package leestcode_0002_Add_Two_Numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	next := 0
	for l1 != nil || l2 != nil {
		var v1, v2 = 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		cur.Val = next + v1 + v2
		next = 0
		if cur.Val >= 10 {
			cur.Val = cur.Val % 10
			next = 1
		}
		if l1 != nil || l2 != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	if next > 0 {
		cur.Next = &ListNode{
			Val: next,
		}
	}

	return head
}