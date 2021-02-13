package Check_If_N_and_Its_Double_Exist

func checkIfExist(arr []int) bool {
	mHash := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if arr[i] % 2 == 0 {
			mHash[arr[i] / 2] = i
		}
	}
	for i := 0; i < len(arr); i++ {
		if j, ok := mHash[arr[i]]; ok && i != j {
			return true
		}
	}
	return false
}