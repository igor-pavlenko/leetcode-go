package leetcode_0141_Linked_List_Cycle

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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	h := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := h[head]; ok {
			return true
		}
		h[head] = struct{}{}
		head = head.Next
	}
	return false
}

// Floyd's Cycle Finding Algorithm
func hasCycleFloyd(head *ListNode) bool {
	if head == nil {
		return false
	}
	s := head
	f := head.Next
	for s != f {
		if f == nil || f.Next == nil {
			return false
		}
		s = s.Next
		f = f.Next.Next
	}
	return true
}