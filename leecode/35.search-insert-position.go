package leecode

// [1,3,5,6]   	5  =>  2
// [1,3,5,6]   	2  =>  1
// [1,3,5,6]   	7  =>  4

func SearchInsert(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	var left = 0
	var right = len(nums)

	for left < right {
		var middle = left + (right-left)/2
		if nums[middle] == target {
			return middle
		} else if target < nums[middle] {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return left
}
