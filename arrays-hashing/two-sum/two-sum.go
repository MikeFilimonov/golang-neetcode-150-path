package twosum

func twoSum(nums []int, target int) []int {

	buffer := make(map[int]int)
	for k, v := range nums {

		another := target - v
		if idx2, found := buffer[another]; found {
			return []int{idx2, k}
		}
		buffer[v] = k

	}

	return []int{}

}
