func maxArea(heights []int) int {

	l,r := 0, len(heights)-1

    maxArea := 0
	for l<r{
		width := r-l
		height:= min (heights[l], heights[r])
		area := width * height
		maxArea = max(area, maxArea)

		if (heights[l] > heights[r]) {
			r --
		}else if (heights[l] < heights[r]){
			l ++
		}else{
			r--
			l++
		}
	}

	return maxArea

}

func min(a, b int) int{
	if a < b{
		return a
	}
	return b
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}