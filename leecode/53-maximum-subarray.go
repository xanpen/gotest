package leecode

// [-2,1,-3,4,-1,2,1,-5,4]
// f(i) = max( num[i], f(i-1) + num[i] )
// f(i)表示第i位及其之前的i-1位组成的所有连续的组合中家和最大的那个值
// 比如:
// i=0  f(0)=-2  只有一种情况：-2
// i=1  f(1)=1  因为f(0)是个负数，所有f(1)中只要num[1]参与进来必然会变小，所以加和最大就是num[1],除非f(0)为正数
// i=2  f(2)=-2  同理，但这个时候f(1)是正数，所以需要num[2]+f(1)=-3+1=-2
func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	var iSum = make(map[int]int)
	iSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		curSum := iSum[i-1] + nums[i]
		if nums[i] >= curSum {
			iSum[i] = nums[i]
		} else {
			iSum[i] = curSum
		}
	}

	const IntMax = int(^uint(0) >> 1)
	const IntMin = ^IntMax

	res := IntMin
	for _, sum := range iSum {
		if sum >= res {
			res = sum
		}
	}

	return res
}
