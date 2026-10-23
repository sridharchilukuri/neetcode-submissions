func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	l, maxCount, res := 0, 0, 0

	for r := 0; r < len(s); r++ {
		count[s[r]]++
		maxCount = max(maxCount, count[s[r]])

		for r-l+1-maxCount > k {
			count[s[l]]--
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}