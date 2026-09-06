package validanagram

func isAnagram(s string, t string) bool {

	const abcLength = 26

	if len(s) != len(t) {
		return false
	}

	var freq [abcLength]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	return freq == [abcLength]int{}

}
