func minWindow(s string, t string) string {
    if len(t) == 0 || len(t) > len(s) {
		return ""
	}

	tmap := make(map[byte]int)

	for i:= 0 ; i < len(t) ; i++{
		tmap[t[i]]++
	}

	sCount:= make(map[byte]int)

    best,l:="", 0

	for r:=0; r < len(s) ; r++{
	  sCount[s[r]]++

	  for mapExists(tmap, sCount){
			best = minString(best, s[l:r+1])
			sCount[s[l]]--
            l++
		}
	}

	return best
}


func mapExists(required, window map[byte]int) bool {
	for ch, count := range required {
		if window[ch] < count {
			return false
		}
	}
	return true
}

func minString(best, candidate string) string {
	if best == "" || len(candidate) < len(best) {
		return candidate
	}
	return best
}