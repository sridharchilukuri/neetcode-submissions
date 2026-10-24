type Entry struct{
		value int
		idx int

	}
func dailyTemperatures(temperatures []int) []int {
	stack := []Entry{}
	res := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1].value < temp {
			entry := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[entry.idx] = i - entry.idx
		}

		stack = append(stack, Entry{value: temp, idx: i})
	}
	return res
}
