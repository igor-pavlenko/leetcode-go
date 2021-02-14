package leetcode_0414_Third_Maximum_Number

// []int{1, 1, 2}
// []int{2, 1, 1} map= 2: 0, 1: 1
func thirdMax(nums []int) int {
	set := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = struct{}{}
	}
	firstMax := 0
	setLen := len(set)
	for i := 0; i < setLen; i++ {
		max := -1 << 31
		for v := range set {
			if v > max {
				max = v
			}
		}
		if i == 0 {
			firstMax = max
		} else if i == 2 {
			return max
		}
		delete(set, max)
	}
	return firstMax
}
