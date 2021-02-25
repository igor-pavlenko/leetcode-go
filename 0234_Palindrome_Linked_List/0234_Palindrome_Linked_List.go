package LEETCODE_0234_Palindrome_Linked_List

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
func isPalindrome(head *ListNode) bool {
	var a []int
	cur := head
	for cur != nil {
		a = append(a, cur.Val)
		cur = cur.Next
	}

	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		if a[i] != a[j] {
			return false
		}
	}

	return true
}