package leecode

func ValidPalindrome(s string) bool {
	if s == "" {
		return false
	}

	// "abca"
	var chars = []rune(s)
	var left = 0
	var right = len(chars) - 1
	var flag = 0

	for left <= right {
		if flag > 1 {
			return false
		}

		if chars[left] != chars[right] && chars[left+1] != chars[right] && chars[left] != chars[right-1] {
			return false
		}

		if chars[left] == chars[right] {
			left++
			right--
		} else if chars[left+1] == chars[right] {
			flag++
			left++
		} else if chars[left] == chars[right-1] {
			flag++
			right--
		}
	}
	return true
}
