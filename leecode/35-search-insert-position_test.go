package leecode

import (
	"log"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	// [1,3,5,6]   	5  =>  2
	// [1,3,5,6]   	2  =>  1
	// [1,3,5,6]   	7  =>  4
	var nums = []int{1, 3, 5, 6}
	log.Println(SearchInsert(nums, 5))
	log.Println(SearchInsert(nums, 2))
	log.Println(SearchInsert(nums, 7))
}
