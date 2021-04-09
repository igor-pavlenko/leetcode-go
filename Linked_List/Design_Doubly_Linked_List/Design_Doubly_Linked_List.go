package Design_Doubly_Linked_List

type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.head == nil {
		return -1
	}
	cur := this.head
	for i := 0; cur != nil; i++ {
		if i == index {
			return cur.Val
		}
		cur = cur.Next
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{
		Val:  val,
		Prev: nil,
		Next: nil,
	}
	if this.head == nil {
		this.head = newNode
		return
	}
	newNode.Next = this.head
	this.head.Prev = newNode
	this.head = newNode
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
		return
	}
	cur := this.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{
		Val:  val,
		Prev: cur,
		Next: nil,
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
// 1->2->4->nil
// 0  1  2      << 2, 3
// 1->2->3->4->nil
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	cur := this.head
	for i := 0; cur != nil; i++ {
		if i == index {
			newNode := &ListNode{
				Val:  val,
				Prev: cur.Prev,
				Next: cur,
			}
			cur.Prev.Next = newNode
			cur.Prev = newNode
			return
		}
		cur = cur.Next
	}
	this.AddAtTail(val)
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.head == nil {
		return
	}
	cur := this.head
	for i := 0; cur != nil; i++ {
		if i == index {
			if cur.Prev == nil {
				// head
				if cur.Next == nil {
					// remove head
					this.head = nil
				} else {
					cur.Next.Prev = nil
					this.head = cur.Next
				}
			} else if cur.Next == nil {
				// tail
				if cur.Prev == nil {
					this.head = nil
				} else {
					// tail
					cur.Prev.Next = nil
				}
			} else {
				cur.Prev.Next = cur.Next
				cur.Next.Prev = cur.Prev
			}
			return
		}
		cur = cur.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
