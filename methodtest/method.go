package methodtest

import (
	"fmt"
)

type car struct {
	name  string
	color string
}

func (c *car) run() {
	fmt.Println(c.name, "car run.")
}

func (c *car) start() string {

	return c.name + " started."
}

// MethodTest test
func MethodTest() {

	c1 := new(car)
	c1.name = "Ben chi"
	c1.color = "yellow"
	s := c1.start()
	fmt.Println(s)
	c1.run()
}
