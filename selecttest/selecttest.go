package selecttest

import (
	"fmt"
	"time"
)

func After() {
	fmt.Println("After Test ", time.Now())
	ch := <-time.After(2 * time.Second)
	// below will print after 2 second
	fmt.Println("now is ", ch)

}

func Tick() {
	// tick test
	fmt.Println("before tick ", time.Now())

	c := 1
	for {
		if c > 5 {
			return
		}
		tick2 := <-time.Tick(1 * time.Second)
		fmt.Println("tick2 is ", tick2)
		c++
	}
}

// SelectTest test
func SelectTest() {

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
