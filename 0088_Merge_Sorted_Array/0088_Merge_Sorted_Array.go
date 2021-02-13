package leetcode_0088_Merge_Sorted_Array

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	i, j := 0, 0
	for n > 0 {
		if i >= m || nums2[j] <= nums1[i] {
			for l := m - 1; l >= i; l-- {
				nums1[l+1] = nums1[l]
			}
			nums1[i] = nums2[j]
			m++
			n--
			j++
		}
		if i < m {
			i++
		}
	}
	return nums1
}
