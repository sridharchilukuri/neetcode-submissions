func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break // sorted: nothing after can sum to 0
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		target := -nums[i]

		for l < r {
			sum := nums[l] + nums[r]
			switch {
			case sum < target:
				l++
			case sum > target:
				r--
			default:
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return result
}