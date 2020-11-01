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
