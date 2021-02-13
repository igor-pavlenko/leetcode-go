package leetcode_0941_Valid_Mountain_Array

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	max := arr[0]
	k := 0
	for i := 1; i < len(arr) && arr[i] > max; i++ {
		max = arr[i]
		k++
	}
	min := max
	l := k
	for j := k + 1; j < len(arr) && arr[j] < min; j++ {
		min = arr[j]
		l++
	}
	if k > 0 && k < l && l == len(arr) - 1 {
		return true
	}
	return false
}
