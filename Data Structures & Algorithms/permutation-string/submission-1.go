func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	s1Count:= [26]int{}

	s2Count:= [26]int{}

	for i := 0; i < len(s1); i++ {
		s1Count[s1[i] - 'a']++
	}

	l:=0

	for r := 0 ; r < len(s2); r++{

		s2Count[s2[r] - 'a' ]++

		if r-l+1 > len(s1) {
			s2Count[s2[l] - 'a']--
			l++
		}


		if (r-l+1 == len (s1)){
            if (s1Count == s2Count){
				return true
			}
		}
	}

return false
}
