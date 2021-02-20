package leetcode_0206_Reverse_Linked_List

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
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	cur := head
	next := cur.Next
	for next != nil {
		tmp := next.Next
		cur.Next = prev
		next.Next = cur
		prev = cur
		cur = next
		next = tmp
	}

	return cur
}
