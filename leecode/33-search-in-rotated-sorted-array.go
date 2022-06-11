package leecode

// [33,34,35,36,37,38, 1, 2, 3, 4, 5]
func SearchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	var left = 0
	var right = len(nums) - 1

	for left <= right {
		var mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			if target < nums[left] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
