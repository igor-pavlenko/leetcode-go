package leetcode_1089_Duplicate_Zeros

func duplicateZeros(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			i++
		}
	}
	return arr
}
