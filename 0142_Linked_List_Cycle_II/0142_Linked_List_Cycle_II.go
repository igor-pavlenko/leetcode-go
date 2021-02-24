package leetcode_0142_Linked_List_Cycle_II

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
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	intersection := detectIntersection(head)
	if intersection == nil {
		return nil
	}

	ptr1 := head
	ptr2 := intersection
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1
}

func detectIntersection(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	f, s := head, head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
		if s == f {
			return f
		}
	}
	return nil
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := h[head]; ok {
			return head
		}
		h[head] = struct{}{}
		head = head.Next
	}
	return nil
}