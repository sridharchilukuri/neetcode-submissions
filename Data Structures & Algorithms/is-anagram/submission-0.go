func isAnagram(s string, t string) bool {

	if len(s) != len(t){
		return false
	}

    sMap:= make (map[rune]int)

	for _,r := range s{
         sMap[r]++
	}
	

	for _, v := range t{
		 sMap[v]--
		 if sMap[v] < 0{
			return false
		 }

	}

    return true

}
