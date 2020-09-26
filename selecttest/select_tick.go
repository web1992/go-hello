package selecttest

import (
	"fmt"
	"time"
)

func SelectTick() {
	tick1 := time.Tick(1 * time.Second)
	tick2 := time.Tick(8 * time.Second)
	for {

		select {
		case <-tick1:
			{
				fmt.Println("time tick is 2 ", time.Now().Second())
			}
		case <-tick2:
			{
				fmt.Println("time tick is 8 ", time.Now().Second())
				return
			}

		}
	}
}
