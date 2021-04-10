package leetcode_0415_Add_Strings

func addStrings(num1 string, num2 string) string {
	var toInt = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var toStr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	carry := 0
	var result string
	var n1, n2 int
	for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0; i, j = i - 1, j - 1 {
		if i < 0 {
			n1 = 0
		} else {
			n1 = toInt[num1[i] - '0']
		}
		if j < 0 {
			n2 = 0
		} else {
			n2 = toInt[num2[j] - '0']
		}

		res := n1 + n2 + carry
		carry = res / 10
		res  = res % 10
		result = toStr[res] + result
	}
	if carry > 0 {
		result = "1" + result
	}
	return result
}