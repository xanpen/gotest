package leecode

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var matrix = [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	// b a b a d
	// 0,1  1,2  2,3  3,4

	// 0,2  1,3  2,4
	// 0,3  1,4
	// 0,4

	for d := 2; d < len(matrix); d++ {
		for i := 0; i < len(matrix)-d; i++ {
			j := i + d
			matrix[i][j] = 1
		}
	}

	for i := 0; i < len(matrix); i++ {
		matrix[i][i] = 1
	}

	for i := 0; i < len(matrix)-1; i++ {
		matrix[i][i+1] = 1
	}

	for _, m := range matrix {
		fmt.Println(m)
	}
}
