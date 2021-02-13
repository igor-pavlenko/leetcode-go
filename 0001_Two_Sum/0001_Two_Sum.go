package leetcode_0001_Two_Sum

import (
	"math"
)

func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 || len(nums) >= int(math.Pow(10, 3)) {
		return nil
	}
	if target <= -1*int(math.Pow(10, 9)) || target >= int(math.Pow(10, 9)) {
		return nil
	}

	h := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		key := target - nums[i]
		h[key] = i
	}

	for i := 0; i < len(nums); i++ {
		j, ok := h[nums[i]]
		if ok && i != j {
			return []int{i, j}
		}
	}

	return nil
}
