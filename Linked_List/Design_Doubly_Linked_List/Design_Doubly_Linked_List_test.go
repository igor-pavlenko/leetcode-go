package Design_Doubly_Linked_List

import (
	"fmt"
	"testing"
)

func printList(head *Node) string {
	result := ""
	for head != nil {
		result = fmt.Sprintf("%s->%d", result, head.Val)
		head = head.Next
	}
	return result
}

func TestMyLinked(t *testing.T) {
	MyList := &MyLinkedList{}
	type args struct {
		action string
		val    int
		index  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Step 1",
			args: args{
				action: "AddAtHead",
				val:    3,
				index:  0,
			},
			want: "->3",
		},
		{
			name: "Step 2",
			args: args{
				action: "AddAtHead",
				val:    2,
				index:  0,
			},
			want: "->2->3",
		},
		{
			name: "Step 3",
			args: args{
				action: "AddAtHead",
				val:    1,
				index:  0,
			},
			want: "->1->2->3",
		},
		{
			name: "Step 4",
			args: args{
				action: "AddAtTail",
				val:    4,
				index:  0,
			},
			want: "->1->2->3->4",
		},
		{
			name: "Step 5",
			args: args{
				action: "AddAtTail",
				val:    5,
				index:  0,
			},
			want: "->1->2->3->4->5",
		},
		{
			name: "Step 6",
			args: args{
				action: "AddAtIndex",
				val:    0,
				index:  0,
			},
			want: "->0->1->2->3->4->5",
		},
		{
			name: "Step 7",
			args: args{
				action: "AddAtIndex",
				val:    6,
				index:  5,
			},
			want: "->0->1->2->3->4->6->5",
		},
		{
			name: "Step 8",
			args: args{
				action: "AddAtIndex",
				val:    7,
				index:  2,
			},
			want: "->0->1->7->2->3->4->6->5",
		},
		{
			name: "Step 9",
			args: args{
				action: "DeleteAtIndex",
				val:    0,
				index:  2,
			},
			want: "->0->1->2->3->4->6->5",
		},
		{
			name: "Step 10",
			args: args{
				action: "DeleteAtIndex",
				val:    0,
				index:  0,
			},
			want: "->1->2->3->4->6->5",
		},
		{
			name: "Step 11",
			args: args{
				action: "DeleteAtIndex",
				val:    0,
				index:  5,
			},
			want: "->1->2->3->4->6",
		},
		{
			name: "Step 12",
			args: args{
				action: "Get",
				val:    0,
				index:  0,
			},
			want: "1",
		},
		{
			name: "Step 13",
			args: args{
				action: "Get",
				val:    0,
				index:  4,
			},
			want: "6",
		},
		{
			name: "Step 14",
			args: args{
				action: "Get",
				val:    0,
				index:  10,
			},
			want: "-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if tt.args.action == "AddAtHead" {
				MyList.AddAtHead(tt.args.val)
			} else if tt.args.action == "AddAtTail" {
				MyList.AddAtTail(tt.args.val)
			} else if tt.args.action == "AddAtIndex" {
				MyList.AddAtIndex(tt.args.index, tt.args.val)
			} else if tt.args.action == "DeleteAtIndex" {
				MyList.DeleteAtIndex(tt.args.index)
			} else if tt.args.action == "Get" {
				gotVal := MyList.Get(tt.args.index)
				got = fmt.Sprintf("%d", gotVal)
				if got != tt.want {
					t.Errorf("%s %s got = %s, want %v", tt.name, tt.args.action, got, tt.want)
				}
				return
			}
			got = printList(MyList.head)
			if got != tt.want {
				t.Errorf("%s %s got = %s, want %v", tt.name, tt.args.action, got, tt.want)
			}
		})
	}
}

//["addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
//[[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]
func TestMyLinked2(t *testing.T) {
	MyList := &MyLinkedList{}
	type args struct {
		action string
		val    int
		index  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Step 1",
			args: args{
				action: "AddAtHead",
				val:    7,
				index:  0,
			},
			want: "->7",
		},
		{
			name: "Step 2",
			args: args{
				action: "AddAtHead",
				val:    2,
				index:  0,
			},
			want: "->2->7",
		},
		{
			name: "Step 3",
			args: args{
				action: "AddAtHead",
				val:    1,
				index:  0,
			},
			want: "->1->2->7",
		},
		{
			name: "Step 4",
			args: args{
				action: "AddAtIndex",
				val:    0,
				index:  3,
			},
			want: "->1->2->7->0",
		},
		//{
		//	name: "Step 5",
		//	args: args{
		//		action: "AddAtTail",
		//		val:    5,
		//		index:  0,
		//	},
		//	want: "->1->2->3->4->5",
		//},
		//{
		//	name: "Step 6",
		//	args: args{
		//		action: "AddAtIndex",
		//		val:    0,
		//		index:  0,
		//	},
		//	want: "->0->1->2->3->4->5",
		//},
		//{
		//	name: "Step 7",
		//	args: args{
		//		action: "AddAtIndex",
		//		val:    6,
		//		index:  5,
		//	},
		//	want: "->0->1->2->3->4->6->5",
		//},
		//{
		//	name: "Step 8",
		//	args: args{
		//		action: "AddAtIndex",
		//		val:    7,
		//		index:  2,
		//	},
		//	want: "->0->1->7->2->3->4->6->5",
		//},
		//{
		//	name: "Step 9",
		//	args: args{
		//		action: "DeleteAtIndex",
		//		val:    0,
		//		index:  2,
		//	},
		//	want: "->0->1->2->3->4->6->5",
		//},
		//{
		//	name: "Step 10",
		//	args: args{
		//		action: "DeleteAtIndex",
		//		val:    0,
		//		index:  0,
		//	},
		//	want: "->1->2->3->4->6->5",
		//},
		//{
		//	name: "Step 11",
		//	args: args{
		//		action: "DeleteAtIndex",
		//		val:    0,
		//		index:  5,
		//	},
		//	want: "->1->2->3->4->6",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if tt.args.action == "AddAtHead" {
				MyList.AddAtHead(tt.args.val)
			} else if tt.args.action == "AddAtTail" {
				MyList.AddAtTail(tt.args.val)
			} else if tt.args.action == "AddAtIndex" {
				MyList.AddAtIndex(tt.args.index, tt.args.val)
			} else if tt.args.action == "DeleteAtIndex" {
				MyList.DeleteAtIndex(tt.args.index)
			} else if tt.args.action == "Get" {
				gotVal := MyList.Get(tt.args.index)
				got = fmt.Sprintf("%d", gotVal)
				if got != tt.want {
					t.Errorf("%s %s got = %s, want %v", tt.name, tt.args.action, got, tt.want)
				}
				return
			}
			got = printList(MyList.head)
			if got != tt.want {
				t.Errorf("%s %s got = %s, want %v", tt.name, tt.args.action, got, tt.want)
			}
		})
	}
}