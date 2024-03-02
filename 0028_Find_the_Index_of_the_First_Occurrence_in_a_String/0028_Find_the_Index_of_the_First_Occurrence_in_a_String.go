package leetcode_0028_Find_the_Index_of_the_First_Occurrence_in_a_String

func strStr(haystack string, needle string) int {
	if len(needle) >= len(haystack) && needle != haystack {
		return -1
	}

	haystack_len := len(haystack)
	needle_len := len(needle)
	res := -1
	i := 0
	j := 0
	for i < haystack_len {
		if haystack[i] == needle[j] {
			if res == -1 {
				res = i
			}
			i++
			j++
			if j == needle_len {
				return res
			}
		} else if res >= 0 {
			i = res + 1
			j = 0
			res = -1
		} else {
			i++
			j = 0
		}
	}
	if j < needle_len {
		return -1
	}
	return res
}
