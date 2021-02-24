package leetcode_0019_Remove_Nth_Node_From_End_of_List

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var list []*ListNode
	cur := head
	for cur != nil {
		list = append(list, cur)
		cur = cur.Next
	}
	if len(list) == 1 {
		return nil
	}
	prevIdx := len(list) - 1 - n
	if prevIdx == -1 {
		head = head.Next
	} else {
		prev := list[prevIdx]
		prev.Next = prev.Next.Next
	}
	return head
}