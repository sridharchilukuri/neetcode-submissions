func lengthOfLongestSubstring(s string) int {

    if len(s)==0{
		return 0
	}

    lastSeen := make(map[byte]int) 
	
	best, left := 0,0

    for right:=0 ; right < len(s); right++{
  
          if idx,ok := lastSeen[s[right]];ok && idx >= left{
                left = idx+1
		  }
		  lastSeen[s[right]] = right

		  best = max(best , right-left+1)
	}

	return best

}