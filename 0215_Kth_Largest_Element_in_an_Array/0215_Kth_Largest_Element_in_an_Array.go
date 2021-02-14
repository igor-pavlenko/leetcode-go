package leetcode_0215_Kth_Largest_Element_in_an_Array

func findKthLargest(nums []int, k int) int {
	max := 0
	for i := 0; i < k; i++ {
		curMax := nums[i]
		curMaxIds := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > curMax {
				curMax = nums[j]
				curMaxIds = j
			}
		}
		tmp := nums[i]
		nums[i] = curMax
		nums[curMaxIds] = tmp
		max = curMax
	}
	return max
}