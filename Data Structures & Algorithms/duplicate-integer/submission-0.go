func hasDuplicate(nums []int) bool {

	seen := make(map[int] bool)


	for _,n := range nums{
       
	   if _, exists := seen[n]; exists{
		return true
	   }

	   seen[n] = true
	}

	return false
}
