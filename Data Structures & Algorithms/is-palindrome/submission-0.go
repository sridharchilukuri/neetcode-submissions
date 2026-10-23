func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !isAlnum(s[l]) {
			l++
			continue
		}
		if !isAlnum(s[r]) {
			r--
			continue
		}
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}


func toLower(c byte) byte {
	if isUpper(c) {
		return c + ('a' - 'A')
	}
	return c
}

func isLower(c byte) bool { return c >= 'a' && c <= 'z' }
func isUpper(c byte) bool { return c >= 'A' && c <= 'Z' }
func isDigit(c byte) bool { return c >= '0' && c <= '9' }

func isAlnum(c byte) bool {
	return isLower(c) || isUpper(c) || isDigit(c)
}