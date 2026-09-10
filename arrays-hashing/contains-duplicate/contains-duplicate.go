package containsduplicate

func hasDuplicate(nums []int) bool {

	if len(nums) < 2 {
		return false
	}

	cloneBuster := make(map[int]struct{})
	for _, v := range nums {

		if _, found := cloneBuster[v]; found {
			return true
		}
		cloneBuster[v] = struct{}{}
	}

	return false
}
