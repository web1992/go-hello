package stacktest

import (
	"fmt"
	"testing"
)

func TestMyStack(t *testing.T) {
	stack := NewStack()

	stack.push("a")
	stack.push("b")
	stack.push("c")
	//stack.log()
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
