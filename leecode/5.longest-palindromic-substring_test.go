package leecode

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	var strs = []string{
		"babad",
		"cbbd",
		"",
		"cbbdccc",
		"aa",
	}

	for _, str := range strs {
		longestPalindrome := LongestPalindrome(str)
		log.Printf("源字符串：%s 的最长子串：%s\n", str, longestPalindrome)
	}
}
