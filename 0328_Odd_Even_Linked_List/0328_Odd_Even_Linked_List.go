package leetcode_0328_Odd_Even_Linked_List

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
func oddEvenList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var headOdd, headEven, curOdd, curEven *ListNode
	i := 1
	for head != nil {
		if i % 2 == 0 {
			if headEven == nil {
				headEven = head
			} else {
				curEven.Next = head
			}
			curEven = head
		} else {
			if headOdd == nil {
				headOdd = head
			} else {
				curOdd.Next = head
			}
			curOdd = head
		}
		head = head.Next
		i++
	}
	if curEven != nil {
		curEven.Next = nil
	}
	if curOdd != nil {
		curOdd.Next = headEven
	}
	return headOdd
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even := head.Next
	headEven := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = headEven
	return head
}
