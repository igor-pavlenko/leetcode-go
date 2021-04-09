package leetcode_0243_Shortest_Word_Distance

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	min, w1Pos, w2Pos := 1<<31, -1, -1
	for i, v := range wordsDict {
		if v == word1 {
			w1Pos = i
		} else if v == word2 {
			w2Pos = i
		}
		if w1Pos >= 0 && w2Pos >= 0 {
			localMin := w2Pos - w1Pos
			if localMin < 0 {
				localMin = -localMin
			}
			if localMin < min {
				min = localMin
			}
		}
	}

	return min
}
