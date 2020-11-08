package stacktest

import (
	"fmt"
	"testing"
)

func TestStack1(t *testing.T) {

	stack := New()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
