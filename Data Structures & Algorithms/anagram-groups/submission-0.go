func groupAnagrams(strs []string) [][]string {

    groups := make(map[[26]int][]string)
        
    for _, s := range strs {
		 counter := [26]int{}
		for _,ch := range s {
			counter[ch-'a']++
		}

		groups[counter] = append(groups[counter], s)
	}


	results:= make ([][]string, 0)

	for _, group:= range groups{
          results = append( results, group)
	}

	return results

}
