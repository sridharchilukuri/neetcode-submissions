func twoSum(nums []int, target int) []int {

	dic := make (map[int]int)

	for i, n := range nums{
        diff := target - n // we get rest

		if value,exists := dic[diff] ; exists{
			return []int{value, i}
		}

		dic[n] = i
	}

	return []int{}
}
