package queuetest

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {

	queue := NewQueue()

	queue.Enter(1)
	queue.Enter(2)
	queue.Enter(3)

	fmt.Println(queue.Out())
	fmt.Println(queue.Out())
	fmt.Println(queue.Out())
	if ok, v := queue.Out(); ok {
		fmt.Println("val is", v)
	} else {
		fmt.Println("queue is empty")
	}
}
