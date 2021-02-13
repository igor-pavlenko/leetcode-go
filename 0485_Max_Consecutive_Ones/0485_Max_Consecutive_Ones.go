package leetcode_0485_Max_Consecutive_Ones

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	curmax := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			curmax++
		} else {
			if curmax >= max {
				max = curmax
			}
			curmax = 0
		}
	}
	if curmax >= max {
		max = curmax
	}
	return max
}
