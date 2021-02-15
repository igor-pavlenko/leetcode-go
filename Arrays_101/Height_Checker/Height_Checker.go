package Height_Checker

func heightChecker(heights []int) int {
	target := make([]int, len(heights))
	copy(target, heights)
	for i := 0; i < len(target); i++ {
		min := target[i]
		minIdx := i
		for j := i + 1; j < len(target); j++ {
			if target[j] < min {
				min = target[j]
				minIdx = j
			}
		}
		tmp := target[i]
		target[i] = target[minIdx]
		target[minIdx] = tmp
	}

	res := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != target[i] {
			res++
		}
	}

	return res
}