func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	s1Count:= make(map[byte]int)

	s2Count:= make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]]++
	}

	l:=0

	for r := 0 ; r < len(s2); r++{

		s2Count[s2[r]]++

		if r-l+1 > len(s1) {
			s2Count[s2[l]]--
			l++
		}


		if (r-l+1 == len (s1)){
            included := true
            for k, _ := range s1Count{
				if s1Count[k] != s2Count[k]{
					included = false
					break
				}
			}
			if included {
				return true
			}
		}
	}

return false
}
