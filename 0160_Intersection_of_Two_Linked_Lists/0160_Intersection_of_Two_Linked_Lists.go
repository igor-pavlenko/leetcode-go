package leetcode_0160_Intersection_of_Two_Linked_Lists

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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	h := make(map[*ListNode]struct{})
	for {
		if headA != nil {
			if _, okA := h[headA]; okA {
				return headA
			}
			h[headA] = struct{}{}
			headA = headA.Next
		}
		if headB != nil {
			if _, okB := h[headB]; okB {
				return headB
			}
			h[headB] = struct{}{}
			headB = headB.Next
		}
		if headA == nil && headB == nil {
			break
		}
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
