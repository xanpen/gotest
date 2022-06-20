package leecode

func ReverseInt(x int) int {
	var rev int32 = 0

	// 0x00000000  取反得到 0xffffffff
	// 右移一位得到 0x7fffffff
	// 最后类型转换得到32为int最大值
	const MaxInt32 = int32(^uint32(0) >> 1)
	const MinInt32 = ^MaxInt32
	for x != 0 {
		remainder := int32(x % 10)

		// MaxInt32 = 2147483647
		// 		      214748365
		// 观察上面的数字，如果翻转后的数字大于214748364，比如：214748365. 那不管余数是什么，必然大于最大值了，溢出了.
		// 如果刚好等于214748364，那就看余数是不是大于7
		// 这里还可以优化：因为题目提到说：给你一个32位整数作为输入，所以理论上不会出现
		// 所以 (rev == MaxInt32 && remainder > 7) 是可以去掉的
		if rev > MaxInt32/10 || (rev == MaxInt32 && remainder > 7) {
			return 0
		}

		// MinInt32 = -2147483648
		// 			  -214748365
		// 同理
		if rev < MinInt32/10 || (rev == MaxInt32 && remainder < -8) {
			return 0
		}

		rev = rev*10 + remainder
		x /= 10
	}

	return int(rev)
}
