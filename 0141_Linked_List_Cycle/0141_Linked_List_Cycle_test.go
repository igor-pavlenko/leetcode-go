package leetcode_0141_Linked_List_Cycle

import "testing"

func TestHasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: false,
		},
		{
			name: "Example 2",
			args: args{},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
