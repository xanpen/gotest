package leecode

func LongestPalindrome(s string) string {
	var bytes = []byte(s)
	var l = len(bytes)
	if l == 0 {
		return ""
	}

	if l == 1 {
		return s
	}

	// Base case: 基础case
	// 		仅包含1个字符的子串肯定为回文串, 如：a b c ...
	// 		仅包含2个字符的子串若这两个字符相同肯定为回文串，如：aa bb cc ...
	// Recursive relation: 递推关系
	// 		包含大于等于3个字符的子串为回文串的条件为：
	// 		收尾字符相同 && 中间剩余的子串也是回文，则整个子串为回文，如 a###a b####b 其中多个#组成的子串也必须是回文

	// 记录最长回文串
	var longestPalindromicArr = bytes[0:1]
	var longestPalindromicLen = 1

	// 使用一个matrix的右上部分来存储记录 可能的子串 是否是回文
	// 设i,j分别为 这个可能的子串的 起始索引和结束索引，也就是：bytes[i] 到 bytes[j] 所代表的子串
	// 如:
	// b a b a d
	// 0,1  1,2  2,3  3,4
	// bytes[0]~bytes[1] 代表ba、bytes[1]~bytes[2] 代表ab ...
	// 同时matrix[i][j]的值 表示 bytes[i] ~ bytes[j] 的串是否为回文
	// 比如：matrix[0][0] = true, 因为 bytes[0]~bytes[0]代表的串是b, 前面已说过，一个字符肯定是回文
	// 这里有些巧妙
	// 有一点说明下，因为我们关注的是从i到j的子串，所以j>=i; 也就是只考虑矩阵的对角线的右侧即可。
	// 也就是下面这些i,j
	// 0,0  1,1  2,2  3,3  4,4
	// 0,1  1,2  2,3  3,4
	// 0,2  1,3  2,4
	// 0,3  1,4
	// 0,4
	var matrix = make([][]int, l)
	for i := 0; i < l; i++ {
		matrix[i] = make([]int, l)
	}

	// 接下来就是根据地推关系设置这些i,j的值 1 or 0
	// 0,0  1,1  2,2  3,3  4,4
	for i := 0; i < l; i++ {
		matrix[i][i] = 1
	}
	// 0,1  1,2  2,3  3,4
	for i := 0; i < l-1; i++ {
		j := i + 1
		if bytes[i] == bytes[j] {
			matrix[i][j] = 1

			if len(bytes[i:j+1]) > longestPalindromicLen {
				longestPalindromicLen = 2
				longestPalindromicArr = bytes[i : j+1]
			}
		}
	}
	// 0,2  1,3  2,4
	// 0,3  1,4
	// 0,4
	for delta := 2; delta < l; delta++ {
		for i := 0; i < l-delta; i++ {
			j := i + delta
			if bytes[i] == bytes[j] && matrix[i+1][j-1] == 1 {
				matrix[i][j] = 1

				if len(bytes[i:j+1]) > longestPalindromicLen {
					longestPalindromicLen = len(bytes[i : j+1])
					longestPalindromicArr = bytes[i : j+1]
				}
			}
		}
	}

	return string(longestPalindromicArr)
}
