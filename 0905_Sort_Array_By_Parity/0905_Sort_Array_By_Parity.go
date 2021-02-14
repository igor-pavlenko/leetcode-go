package leetcode_0905_Sort_Array_By_Parity

// [3,1,2,4]
func sortArrayByParity(A []int) []int {
	writePointer := 0
	for readPointer := 0; readPointer < len(A); readPointer++ {
		if A[readPointer] % 2 == 0 {
			tmp := A[writePointer]
			A[writePointer] = A[readPointer]
			A[readPointer] = tmp
			writePointer++
		}
	}
	return A
}
