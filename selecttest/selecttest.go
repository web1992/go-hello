package selecttest

import (
	"fmt"
	"time"
)

// SelectTest test
func SelectTest() {

	fmt.Println("SelectTest ", time.Now())
	ch := <-time.After(2 * time.Second)
	// below will print after 2 second
	fmt.Println("now is ", ch)

	// select test
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

	// tick test
	fmt.Println("before tick ", time.Now())
	tick2 := <-time.Tick(2 * time.Second)
	fmt.Println("tick2 is ", tick2)

	// for {

	// 	tick1 := time.Tick(1 * time.Second)
	// 	tick2 := time.Tick(8 * time.Second)

	// 	select {
	// 	case <-tick1:
	// 		{
	// 			fmt.Println("time tick is 2 ", time.Now().Second())
	// 		}
	// 	case <-tick2:
	// 		{
	// 			fmt.Println("time tick is 8 ", time.Now().Second())
	// 			return
	// 		}

	// 	}
	// }

	tick := time.Tick(1 * time.Second)
	after := time.After(7 * time.Second)
	fmt.Println("start second:", time.Now().Second())
	for {
		select {
		case <-tick:
			fmt.Println("1 second over:", time.Now().Second())
		case <-after:
			fmt.Println("7 second over:", time.Now().Second())
			return
		}
	}

}
