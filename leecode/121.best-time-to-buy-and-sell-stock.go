package leecode

// [7,1,5,3,6,4]
// 5

// [7,6,4,3,1]
// 0

// [2,4,1]
// 2

// [1]
// 0

func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	const IntMax = int(^uint(0) >> 1)
	const IntMin = ^IntMax

	var minValue = prices[0]
	var maxDiff = IntMin
	for i := 1; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		}
		diff := prices[i] - minValue
		if diff > maxDiff {
			maxDiff = diff
		}
	}

	return maxDiff
}
