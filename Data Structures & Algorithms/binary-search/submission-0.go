func search(nums []int, target int) int {


	var bSearch func(left, right int) int

	bSearch = func(left, right int) int{
		if (left > right){
				return -1
		}

	   	mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			return bSearch(left, mid-1)
		}else{
			return bSearch(mid+1,right)
		}
	}

	return bSearch(0, len(nums)-1)

}
