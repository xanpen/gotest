package leecode

// abcabcbb
func LengthOfLongestSubstring(s string) int {
	var bytes = []byte(s)

	if len(bytes) <= 0 {
		return len(bytes)
	}

	// 用来存放截止到目前出现的字符
	// 用于当右指针向右拓展时指向的字符是否重复，也就是是否出现在map中
	var uniqueMap = make(map[byte]int)
	var left = 0
	var right = 0
	var maxLength = 0

	// 右指针不能越界
	for right < len(bytes) {
		// 向右拓展右指针,直到碰到重复的字符
		for right < len(bytes) {
			// 右指针所指字符重复，退出loop
			if _, ok := uniqueMap[bytes[right]]; ok {
				break
			}
			// 设置用于判重的map
			uniqueMap[bytes[right]] = 1
			// 记录当前子串长度
			if len(bytes[left:right+1]) > maxLength {
				maxLength = len(bytes[left : right+1])
			}
			// 继续尝试拓展
			right++
		}

		// 因为右指针拓展失败了，这个时候需要左指针向右移动一格
		// 需要在判重map中删除左指针在移动之前指向的字符
		delete(uniqueMap, bytes[left])
		left++
	}

	return maxLength
}
