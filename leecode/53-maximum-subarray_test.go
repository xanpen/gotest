package leecode

import (
	"log"
	"testing"
)

// [-2,1,-3,4,-1,2,1,-5,4]
func TestMaxSubArray(t *testing.T) {
	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	sum := maxSubArray(nums)
	log.Println(sum)
}
