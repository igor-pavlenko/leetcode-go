package leetcode_0383_Ransom_Note

func canConstruct(ransomNote string, magazine string) bool {
	mHash := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		cnt := mHash[magazine[i]]
		mHash[magazine[i]] = cnt + 1
	}

	for i := 0; i < len(ransomNote); i++ {
		cnt, ok := mHash[ransomNote[i]]
		if !ok || cnt <= 0 {
			return false
		}
		mHash[ransomNote[i]] = cnt - 1
	}

	return true
}