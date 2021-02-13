package leetcode_0009_Palindrome_Number

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example 1", args: args{x: 121}, want: true},
		{name: "Example 2", args: args{x: -121}, want: false},
		{name: "Example 3", args: args{x: 10}, want: false},
		{name: "Example 4", args: args{x: -101}, want: false},
		{name: "Example 5", args: args{x: 100001}, want: true},
		{name: "Example 6", args: args{x: 5}, want: true},
		{name: "Example 7", args: args{x: 12344321}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
