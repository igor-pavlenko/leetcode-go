package leetcode_0021_Merge_Two_Sorted_Lists

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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val == l2.Val {
				cur.Val = l1.Val
				cur.Next = &ListNode{}
				cur = cur.Next
				cur.Val = l2.Val
				l1 = l1.Next
				l2 = l2.Next
			} else if l1.Val < l2.Val {
				cur.Val = l1.Val
				l1 = l1.Next
			} else if l1.Val > l2.Val {
				cur.Val = l2.Val
				l2 = l2.Next
			}
		} else if l1 != nil {
			cur.Val = l1.Val
			l1 = l1.Next
		} else if l2 != nil {
			cur.Val = l2.Val
			l2 = l2.Next
		}
		if l1 != nil || l2 != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}

	return head
}
