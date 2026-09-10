package groupanagrams

func groupAnagrams(strs []string) [][]string {

	const abcLength = 26

	buffer := make(map[[abcLength]int][]string)

	for _, v := range strs {

		var freq [abcLength]int
		for i := 0; i < len(v); i++ {
			freq[v[i]-'a']++
		}

		buffer[freq] = append(buffer[freq], v)

	}

	result := make([][]string, len(buffer))
	for _, v := range buffer {
		result = append(result, v)
	}

	return result

}
