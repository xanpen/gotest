package leecode

import (
	"log"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var strs = []string{"abcabcbb", "bbbbb", "pwwkew"}
	for _, str := range strs {
		substring := LengthOfLongestSubstring(str)
		log.Printf("源字符串：%s 的 最长不重复子串长度为：%d\n", str, substring)
	}
}
