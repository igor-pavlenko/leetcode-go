package leetcode_0977_Squares_of_a_Sorted_Array

import "math"

func sortedSquares(nums []int) []int {
	var sq []int
	for i := 0; i < len(nums); i++ {
		sq = append(sq, int(math.Pow(float64(nums[i]), 2)))
	}

	for i := 0; i < len(sq)-1; i++ {
		minIndx := i
		for j := i + 1; j < len(sq); j++ {
			if sq[j] < sq[minIndx] {
				minIndx = j
			}
		}

		tmp := sq[i]
		sq[i] = sq[minIndx]
		sq[minIndx] = tmp
	}

	return sq
}
