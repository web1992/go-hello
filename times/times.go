package times

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func testAfter() {
	var wg sync.WaitGroup
	wg.Add(1)

	time.AfterFunc(time.Second*3, func() {
		fmt.Println("done")
		wg.Done()
	})

	wg.Wait()
}

func testContext() {
	ctx := context.Background()

	ctx, cancelFunc := context.WithCancel(ctx)

	time.AfterFunc(time.Second*3, cancelFunc)
}
