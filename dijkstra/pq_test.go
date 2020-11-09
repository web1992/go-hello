package dijkstra

import (
	"fmt"
	"testing"
)

func TestPQ(t *testing.T) {

	pq := &PQ{}
	e1 := &entry{key: "2", priority: 5}
	e2 := &entry{key: "3", priority: 24}
	e3 := &entry{key: "4", priority: 33}
	pq.Push(&entry{key: "5", priority: 4})
	pq.Push(e1)
	pq.Push(e2)
	pq.Push(e3)

	fmt.Println(pq.ToString())
}
