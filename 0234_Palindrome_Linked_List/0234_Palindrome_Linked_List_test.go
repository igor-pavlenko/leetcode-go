package LEETCODE_0234_Palindrome_Linked_List

import "testing"

func TestIsPalindrome(t *testing.T) {
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
					Val:  0,
					Next: &ListNode{
						Val:  1,
						Next: &ListNode{
							Val:  1,
							Next: &ListNode{
								Val:  0,
								Next: nil,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				head: &ListNode{
					Val:  0,
					Next: &ListNode{
						Val:  1,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{
								Val:  0,
								Next: nil,
							},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				head: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: true,
		},
		{
			name: "Example 4",
			args: args{
				head: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
