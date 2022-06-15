package leecode

import (
	"log"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var cases = []int{
		121, -121, 10,
	}

	for _, c := range cases {
		log.Printf("%d 是否为回文数字：%t", c, IsPalindrome(c))
	}
}
