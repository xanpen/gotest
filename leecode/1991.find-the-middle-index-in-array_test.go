package leecode

import (
	"log"
	"testing"
)

func TestFindMiddleIndex(t *testing.T) {
	var numss = [][]int{
		{2, 3, -1, 8, 4}, // 3
		{1, -1, 4},       // 2
		{2, 5},           // -1
		{1},              // 0
	}

	for _, nums := range numss {
		midIndex := FindMiddleIndex(nums)
		log.Println(midIndex)
	}
}
