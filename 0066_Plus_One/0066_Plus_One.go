package leetcode_0066_Plus_One

func plusOne(digits []int) []int {
	inc := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + inc
		if digits[i] == 10 {
			inc = 1
			digits[i] = 0
		} else {
			inc = 0
		}
	}
	if inc > 0 {
		digits = append([]int{inc}, digits...)
	}

	return digits
}
