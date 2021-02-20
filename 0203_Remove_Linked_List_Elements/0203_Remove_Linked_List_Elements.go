package leetcode_0203_Remove_Linked_List_Elements

type ListNode struct {
    Val int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	cur := head
	next := cur.Next
	for next != nil {
		if cur.Val == val {
			if prev == nil {
				cur = next
				next = next.Next
				head = cur
			} else {
				prev.Next = next
				cur = next
				next = next.Next
			}
		} else {
			prev = cur
			cur = next
			next = next.Next
		}
	}
	if cur.Val == val {
		if prev == nil {
			cur = next
			head = cur
		} else {
			prev.Next = next
		}
	}

	return head
}
