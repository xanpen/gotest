package leecode

func IsPalindrome(x int) bool {
	var remainders []int

	if x < 0 {
		return false
	}

	for x > 0 {
		remainder := x % 10
		remainders = append(remainders, remainder)
		x = x / 10
	}

	var left = 0
	var right = len(remainders) - 1

	for left <= right {
		if remainders[left] != remainders[right] {
			return false
		}
		left++
		right--
	}

	return true
}
