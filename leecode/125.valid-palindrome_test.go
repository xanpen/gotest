package leecode

import (
	"fmt"
	"log"
	"testing"
	"unicode"
)

func TestIsPalindromeStr(t *testing.T) {
	var cases = []string{
		"A man, a plan, a canal: Panama",
		"amanaplanacanalpanama",
		"race a car",
		"",
	}

	fmt.Println(unicode.IsLetter(':'))

	for _, s := range cases {
		log.Printf("'%s'是否为回文串：%t", s, IsPalindromeStr(s))
	}
}
