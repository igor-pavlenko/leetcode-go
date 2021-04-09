package leetcode_0067_Add_Binary

func addBinary(a string, b string) string {
	result := ""
	var aInt, bInt, carry int
	carry = 0
	for i, j := len(a) - 1, len(b) - 1; i >= 0 || j >= 0; i, j = i - 1, j - 1 {
		aInt = 0
		if i >= 0 && a[i] == '1' {
			aInt = 1
		}
		bInt = 0
		if j >= 0 && b[i] == '1' {
			bInt = 1
		}
		s := aInt + bInt + carry
		if s == 0 {
			carry = 0
			result = "0" + result
		} else if s == 1 {
			carry = 0
			result = "1" + result
		} else if s == 2 {
			carry = 1
			result = "0" + result
		} else if s == 3 {
			carry = 1
			result = "1" + result
		}
	}
	if carry == 1 {
		result = "1" + result
	}
	return result
}
