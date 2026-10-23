func trap(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax := 0, 0
	total := 0

	for l < r {
		if height[l] < height[r] {
			leftMax = max(leftMax, height[l])
			total += leftMax - height[l]
			l++
		} else {
			rightMax = max(rightMax, height[r])
			total += rightMax - height[r]
			r--
		}
	}
	return total
}