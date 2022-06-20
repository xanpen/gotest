package leecode

func ValidPalindrome(s string) bool {
	if s == "" {
		return false
	}

	// "abca"
	var chars = []rune(s)
	var left = 0
	var right = len(chars) - 1

	var leftFlag = true
	var rightFlag = true

	for left < right {
		if chars[left] == chars[right] {
			left++
			right--
		} else {
			// 多余的那个字符在右侧，跳过他继续检查是不是回文
			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				if chars[i] != chars[j] {
					rightFlag = false
					break
				}
			}
			// 多余的那个字符在左侧， 跳过他继续检查是不是回文
			for i, j := left+1, right; i < j; i, j = i+1, j-1 {
				if chars[i] != chars[j] {
					leftFlag = false
					break
				}
			}
			// 两种情况，只要有一种还是回文返回：是回文
			return leftFlag || rightFlag
		}
	}

	return true
}
