package leecode

import (
	"log"
	"testing"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	var nums = []int{33, 34, 35, 36, 37, 38, 1, 2, 3, 4, 5}
	for _, num := range nums {
		index := SearchInRotatedSortedArray(nums, num)
		log.Println(index)
	}
}
