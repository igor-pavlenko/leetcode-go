package leetcode_0136_Single_Number

func singleNumber(nums []int) int {
	hmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cnt := hmap[nums[i]]
		hmap[nums[i]] = cnt + 1
	}
	for k, v := range hmap {
		if v == 1 {
			return k
		}
	}
	return -1
}
