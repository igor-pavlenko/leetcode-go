package leetcode_0283_Move_Zeroes

// [0,1,0,3,12]
func moveZeroes(nums []int)  {
	writePointer := 0
	for readPointer := 0; readPointer < len(nums); readPointer++ {
		if nums[readPointer] != 0 {
			nums[writePointer] = nums[readPointer]
			writePointer++
		}
	}
	for i := writePointer; i < len(nums); i++ {
		nums[i] = 0
	}
}
