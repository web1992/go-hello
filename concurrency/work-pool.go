package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("work job start", "job id", id)
		results <- j * 2
		fmt.Println("work job en", "job id", id)
	}
}
func workPool(id int, jobs <-chan int, results chan<- int) {

	var wg sync.WaitGroup

	for j := range jobs {
		wg.Add(1)
		go func(jj int) {
			fmt.Println("job start", "job id", id)
			time.Sleep(time.Duration(1) * time.Second)
			fmt.Println("job end", "job id,", id)
			results <- jj * 2
			wg.Done()
		}(j)
	}

	wg.Wait()
}

func main() {
	nums := 8
	jobs := make(chan int, nums)
	result := make(chan int, nums)

	for i := 1; i <= 3; i++ {
		go workPool(i, jobs, result)
	}

	for j := 1; j <= nums; j++ {
		jobs <- j
	}

	close(jobs)
	fmt.Println("Closed job")

	for j := 1; j < nums; j++ {
		<-result
	}
	close(result)
}
