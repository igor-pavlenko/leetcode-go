package Replace_Elements_with_Greatest_Element_on_Right_Side

func replaceElements(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		max := 0
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] > max {
				max = arr[j]
			}
		}
		arr[i] = max
	}
	arr[len(arr) - 1] = -1
	return arr
}