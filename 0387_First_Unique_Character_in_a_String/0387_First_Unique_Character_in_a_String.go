package leetcode_0387_First_Unique_Character_in_a_String

func firstUniqChar(s string) int {
	count := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		j := count[s[i]]
		count[s[i]] = j + 1
	}
	for i := 0; i < len(s); i++ {
		if cnt := count[s[i]]; cnt == 1 {
			return i
		}
	}
	return -1
}
