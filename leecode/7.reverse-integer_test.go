package leecode

import (
	"fmt"
	"testing"
)

func TestReverseInt(t *testing.T) {
	var nums = []int{
		123,
		0,
		002,
		200,
	}
	for _, num := range nums {
		reverseInt := ReverseInt(num)
		fmt.Printf("num %d, reverse Num:%d\n", num, reverseInt)
	}
}
