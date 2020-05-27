package structtest

import (
	"fmt"
	"time"
)

type Position string

type Employee2 struct {
	ID       int
	Name     string
	Address  string
	DoB      time.Time
	Position Position
}

type Point struct {
	X, Y int
}

// String
func (p Point) String() string {
	return fmt.Sprint("x is ", p.X, " Y is ", p.Y)
}

// TestPoint
func TestPoint() {

	point := Point{X: 1, Y: 2}

	pp := &Point{11, 22}
	fmt.Println(point)
	fmt.Println(&pp)
	pp.X = 99
	pp.Y = 100
	fmt.Println(pp)
}
