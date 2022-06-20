package leecode

import (
	"log"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	// [7,1,5,3,6,4]
	// 5

	// [7,6,4,3,1]
	// 0
	var prices = []int{7, 1, 5, 3, 6, 4}
	profit := MaxProfit(prices)
	log.Println(profit)

	var prices2 = []int{7, 6, 4, 3, 1}
	profit2 := MaxProfit(prices2)
	log.Println(profit2)

	var prices3 = []int{2, 4, 1}
	profit3 := MaxProfit(prices3)
	log.Println(profit3)

	var prices4 = []int{3, 2, 6, 5, 0, 3}
	profit4 := MaxProfit(prices4)
	log.Println(profit4)
}
