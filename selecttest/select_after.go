package selecttest

import (
	"fmt"
	"time"
)

func SelectAfter() {
	// select test
	for {
		select {

		case val := <-time.After(2 * time.Second):
			{
				fmt.Println("after 2 second val is ", val)
			}

		case <-time.After(5 * time.Second):
			{
				fmt.Println("after 5 second ")
			}
		}
	}

}
