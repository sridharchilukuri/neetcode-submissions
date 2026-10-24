func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	cols := len(matrix[0])
	left, right := 0, len(matrix)-1


	for left <= right {         
		 mid := left + (right-left)/2

		if target < matrix[mid][0] {
			right = mid - 1
		} else if target > matrix[mid][cols-1] {
			left = mid + 1
		}else{
			lo, hi := 0, cols-1
			for lo <= hi {
				m := lo + (hi-lo)/2
				if matrix[mid][m] == target {
					return true
				}
				if matrix[mid][m] < target {
					lo = m + 1
				} else {
					hi = m - 1
				}
			}
			return false
		}
	}

	return false
}