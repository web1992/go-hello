package factorial

import (
	"fmt"
	"testing"
)

// 一个递归函数
func TestFactorial(t *testing.T) {
	// 5! = 5*4*3*2*1
	i := factorial(5)
	fmt.Println(i)
}

// 使用递归计算 合计
func TestSum(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	s := sum(a)
	fmt.Println(s)
}
