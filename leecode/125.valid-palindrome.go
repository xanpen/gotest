package leecode

import (
	"strings"
	"unicode"
)

// IsPalindromeStr "A man, a plan, a canal: Panama"
func IsPalindromeStr(s string) bool {
	var chars []rune

	if s == "" {
		return true
	}

	for _, r := range strings.ToLower(s) {
		if unicode.IsNumber(r) || unicode.IsLetter(r) {
			chars = append(chars, r)
		}
	}

	var left = 0
	var right = len(chars) - 1

	for left <= right {
		if chars[left] != chars[right] {
			return false
		}

		left++
		right--
	}

	return true
}
