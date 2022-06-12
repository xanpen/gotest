package leecode

// [2,3,-1,8,4] 	3
// [1,-1,4] 		2
// [2,5]			-1
// [1]				0

func FindMiddleIndex(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}

	var sum = 0
	for _, num := range nums {
		sum += num
	}

	// leftSum + rightSum + num[i] = sum
	// leftSum = rightSum
	// => leftSum + leftSum + num[i] = sum
	var leftSum = 0
	for i := 0; i < len(nums); i++ {
		if leftSum+leftSum+nums[i] == sum {
			return i
		} else {
			leftSum += nums[i]
		}
	}

	return -1
}
