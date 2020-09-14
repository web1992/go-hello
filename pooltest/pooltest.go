package pooltest

import (
	"fmt"
	"sync"
)

func testPool() {

	var p sync.Pool

	p.Put("a")

	fmt.Printf("pool get is %s \n", p.Get())
}
