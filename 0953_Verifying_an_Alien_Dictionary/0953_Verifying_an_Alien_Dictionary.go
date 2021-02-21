package leetcode_0953_Verifying_an_Alien_Dictionary

func isAlienSorted(words []string, order string) bool {
	orderHash := make(map[byte]int, 26)
	for i := 0; i < len(order); i++ {
		orderHash[order[i]] = i
	}
	for i := 0; i < len(words) - 1; i++ {
		if compare(words[i], words[i + 1], orderHash) > 0 {
			return false
		}
	}

	return true
}

func compare(w1 string, w2 string, orderHash map[byte]int) int {
	for i := 0; i < len(w1) && i < len(w2); i++ {
		c1 := orderHash[w1[i]]
		c2 := orderHash[w2[i]]
		if c1 < c2 {
			return -1
		} else if c1 > c2 {
			return 1
		}
	}
	if len(w1) < len(w2) {
		return -1
	} else if len(w1) > len(w2) {
		return 1
	}

	return 0
}