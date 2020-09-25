package selecttest

import "time"

func SelectCase() {

	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 1
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		}
	}
}
