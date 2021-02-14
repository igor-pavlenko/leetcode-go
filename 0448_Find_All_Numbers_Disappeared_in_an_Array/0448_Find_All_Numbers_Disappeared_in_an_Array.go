package leetcode_0448_Find_All_Numbers_Disappeared_in_an_Array

func findDisappearedNumbers(nums []int) []int {
	s := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		s[nums[i]] = struct{}{}
	}
	var res []int
	for i := 1; i <= len(nums); i++ {
		if _, ok := s[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}
