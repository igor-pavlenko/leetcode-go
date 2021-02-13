package leetcode_0007_Reverse_Integer

import (
	"math"
)

func reverse(x int) int {
	var result int = 0

	var l []int
	for x != 0 {
		t := x % 10
		x = x / 10
		l = append(l, int(t))
	}
	for i := 0; i < len(l); i++ {
		if result > 1<<31-1 || result < -(1<<31) {
			return 0
		}
		result = result + l[i]*int(math.Pow(10, float64(len(l)-i-1)))
	}
	return result
}
