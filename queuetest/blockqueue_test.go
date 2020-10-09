package queuetest

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBlockQueue(t *testing.T) {

	//var queue *BlockQueue = new(BlockQueue)

	queue := new(BlockQueue)

	var w sync.WaitGroup
	w.Add(1)

	go func() {
		i := 0
		for {
			//fmt.Println("Offer is run")
			i = i + 1
			fmt.Printf("put obj is %d \n", i)
			queue.Offer(i)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			//fmt.Println("Take is run")
			obj := queue.Take()
			fmt.Printf("take obj is %d \n", obj)
			time.Sleep(1 * time.Second)
		}

	}()

	// let main goroutine not exit
	w.Wait()

}
