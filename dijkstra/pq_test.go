package dijkstra

import (
	"fmt"
	"testing"
)

func TestPQPushKey(t *testing.T) {
	pq := &PQ{}
	pq.PushKey("E", 7)
	pq.PushKey("D", 4)
	pq.PushKey("C", 3)
	pq.PushKey("B", 2)
	pq.PushKey("A", 1)

	fmt.Println(pq.PoPKey())
	fmt.Println(pq.PoPKey())
	fmt.Println(pq.PoPKey())
	fmt.Println(pq.PoPKey())
	fmt.Println(pq.PoPKey())

	// Except Out:
	// A 1
	// B 2
	// C 3
	// D 4
	// E 7

}
func TestPQ(t *testing.T) {

	pq := &PQ{}
	e1 := &entry{key: "B", priority: 5}
	e2 := &entry{key: "C", priority: 24}
	e3 := &entry{key: "D", priority: 33}
	pq.Push(&entry{key: "A", priority: 4})
	pq.Push(e1)
	pq.Push(e2)
	pq.Push(e3)
	pq.PushKey("E", 1)

	fmt.Println(pq.PoPKey())
	fmt.Println(pq.PoPKey())
	fmt.Println(pq.PoPKey())
	fmt.Println(pq.PoPKey())
	fmt.Println(pq.PoPKey())

	// Expect Out:
	//E 1
	//A 4
	//B 5
	//C 24
	//D 33
}
