package Design_Linked_List

type ListNode struct {
	val  int
	next *ListNode
}

type MyLinkedList struct {
	head *ListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
	}
}

func (this *MyLinkedList) GetNode(index int) *ListNode {
	if index < 0 {
		return nil
	}
	cur := this.head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.next
	}
	if cur == nil {
		return nil
	}
	return cur
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	cur := this.head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.next
	}
	if cur == nil {
		return -1
	}
	return cur.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	cur := &ListNode{
		val:  val,
		next: this.head,
	}
	this.head = cur
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.head
	for cur != nil && cur.next != nil {
		cur = cur.next
	}
	if cur == nil {
		this.AddAtHead(val)
	} else {
		cur.next = &ListNode{
			val:  val,
			next: nil,
		}
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	cur := this.head
	for i := 0; i < index - 1 && cur != nil; i++ {
		cur = cur.next
	}
	if cur == nil {
		return
	}
	cur.next = &ListNode{
		val:  val,
		next: cur.next,
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.GetNode(index)
	if cur == nil {
		return
	}
	prev := this.GetNode(index - 1)
	next := cur.next
	if prev != nil {
		prev.next = next
	} else {
		this.head = next
	}
}
