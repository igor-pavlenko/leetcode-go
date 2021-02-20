package leetcode_0487_Max_Consecutive_Ones_II

// []int{1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
//       0  1  2  3  4  5  6  7  8  9
// []int{0, 1, 0, 1}
// []int{0, 0, 0, 1}
func findMaxConsecutiveOnes(nums []int) int {

	max := 0
	lastZero := 0
	curMax := 0
	swap := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if !swap {
				swap = true
				curMax++
				if curMax > max {
					max = curMax
				}
				lastZero = i
			} else {
				swap = false
				curMax = i - lastZero - 1
				lastZero = i
				if i > 0 {
					i--
				}
			}
		} else if nums[i] == 1 {
			curMax++
			if curMax > max {
				max = curMax
			}
		}
	}

	return max
}
