package _026_Remove_Duplicates_From_Sorted_Array

// []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4} len = 10
//       0  1  2  3  4  5  6  7  8  9
//
func removeDuplicates(nums []int) []int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		if nums[i] == nums[i+1] {
			l--
			for j := i + 1; j < l; j++ {
				nums[j] = nums[j+1]
			}
			i--
		}
	}
	return nums[0:l]
}

func removeDuplicates2(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return nums[0 : i+1]
}
