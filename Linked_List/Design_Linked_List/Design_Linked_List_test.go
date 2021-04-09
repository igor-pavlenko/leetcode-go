package Design_Linked_List

import (
	"reflect"
	"testing"
)

func TestMyLinkedListAddAtHead(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		list   *MyLinkedList
		name   string
		args   args
	}{
		{
			name: "Example 1",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: nil,
				},
			},
			args: args{val: 2},
		},
		{
			name: "Example 2",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: &ListNode{
						val:  2,
						next: nil,
					},
				},
			},
			args: args{val: 3},
		},
		{
			name: "Example 3",
			list: &MyLinkedList{
				head: nil,
			},
			args: args{val: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.AddAtHead(tt.args.val)
			if tt.list.head != nil && tt.list.head.val != tt.args.val {
				t.Errorf("head.val != %d", tt.args.val)
			}
		})
	}
}

func TestMyLinkedListAddAtTail(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		list   *MyLinkedList
		name   string
		args   args
		want   *MyLinkedList
	}{
		{
			name: "Example 1",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: nil,
				},
			},
			args: args{val: 2},
			want: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: &ListNode{
						val:  2,
						next: nil,
					},
				},
			},
		},
		{
			name: "Example 2",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: &ListNode{
						val:  2,
						next: nil,
					},
				},
			},
			args: args{val: 3},
			want: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: &ListNode{
						val:  2,
						next: &ListNode{
							val:  3,
							next: nil,
						},
					},
				},
			},
		},
		{
			name: "Example 3",
			list: &MyLinkedList{
				head: nil,
			},
			args: args{val: 4},
			want: &MyLinkedList{
				head: &ListNode{
					val:  4,
					next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.AddAtTail(tt.args.val)
			if !reflect.DeepEqual(tt.list, tt.want) {
				t.Errorf("AddAtTail() = %v, want %v", tt.list, tt.want)
			}
		})
	}
}

func TestMyLinkedListGet(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		list   *MyLinkedList
		name   string
		args   args
		want   int
	}{
		{
			name: "Example 1",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: nil,
				},
			},
			args: args{idx: 0},
			want: 1,
		},
		{
			name: "Example 2",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: &ListNode{
						val:  2,
						next: nil,
					},
				},
			},
			args: args{idx: 1},
			want: 2,
		},
		{
			name: "Example 3",
			list: &MyLinkedList{
				head: nil,
			},
			args: args{idx: 1},
			want: -1,
		},
		{
			name: "Example 4",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: nil,
				},
			},
			args: args{idx: 5},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.Get(tt.args.idx)
			if got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyLinkedListAddAtIndex(t *testing.T) {
	type args struct {
		idx int
		val int
	}
	tests := []struct {
		list   *MyLinkedList
		name   string
		args   args
		want   int
	}{
		{
			name: "Example 1",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: nil,
				},
			},
			args: args{idx: 0, val: 2},
			want: 2,
		},
		{
			name: "Example 2",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: &ListNode{
						val:  2,
						next: nil,
					},
				},
			},
			args: args{idx: 1, val: 3},
			want: 3,
		},
		{
			name: "Example 3",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: &ListNode{
						val:  2,
						next: nil,
					},
				},
			},
			args: args{idx: 10, val: 1},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.AddAtIndex(tt.args.idx, tt.args.val)
			got := tt.list.Get(tt.args.idx)
			if got != tt.want {
				t.Errorf("AddAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyLinkedListDeleteAtIndex(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		list   *MyLinkedList
		name   string
		args   args
		want   *MyLinkedList
	}{
		{
			name: "Example 1",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: nil,
				},
			},
			args: args{idx: 0},
			want: &MyLinkedList{
				head: nil,
			},
		},
		{
			name: "Example 2",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: &ListNode{
						val:  2,
						next: nil,
					},
				},
			},
			args: args{idx: 1},
			want: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: nil,
				},
			},
		},
		{
			name: "Example 3",
			list: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: &ListNode{
						val:  2,
						next: &ListNode{
							val:  3,
							next: nil,
						},
					},
				},
			},
			args: args{idx: 1},
			want: &MyLinkedList{
				head: &ListNode{
					val:  1,
					next: &ListNode{
						val:  3,
						next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.DeleteAtIndex(tt.args.idx)
			if !reflect.DeepEqual(tt.list, tt.want) {
				t.Errorf("DeleteAtIndex() = %v, want %v", tt.list, tt.want)
			}
		})
	}
}