package leetcode_0009_Palindrome_Number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	var list []int
	for x != 0 {
		t := x % 10
		x = x / 10
		list = append(list, t)
	}

	for i, j := 0, len(list)-1; i < len(list)/2; i, j = i+1, j-1 {
		if list[i] != list[j] {
			return false
		}
	}
	return true
}
