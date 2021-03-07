package base

import "fmt"

// 表达式估值顺序
func EvaluateOfExpression() {
	var i, arr = 0, [5]int{1, 2, 3, 4, 5}
	i, arr[i] = 2, 99
	fmt.Println(i, arr)
}
