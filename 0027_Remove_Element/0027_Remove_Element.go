package leetcode_0027_Remove_Element

// []int{3, 2, 2, 3}, 3
// j nums
// 0 3, 2, 2, 3 i=0
// 1 2, 2, 2, 3 i=1
// 2 2, 2, 2, 3 i=2
// 3
func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
