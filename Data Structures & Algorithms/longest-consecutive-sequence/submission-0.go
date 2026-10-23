func longestConsecutive(nums []int) int {

	dic := make(map[int]bool)

	for _,n := range nums{
		dic[n] = true
	}
    longest := 0

	for k,_ := range dic{
		if dic[k-1] == true{
			continue //not start of sequence
		}
        
		tmpLongest:= 1

		for dic[k+tmpLongest]{
			tmpLongest++
		}
		
		longest = max(longest, tmpLongest)
	}


	return longest
}


func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
