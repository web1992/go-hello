package dijkstra

import (
	"fmt"
	"testing"
)

func Test_dijkstra(t *testing.T) {
	//dijkstra()
}

func Test(t *testing.T) {

	graph := map[string]map[string]int{
		"A": {"B": 5, "C": 1},
		"B": {"A": 5, "C": 2, "D": 1},
		"C": {"A": 1, "B": 2, "D": 4, "E": 8},
		"D": {"B": 1, "C": 4, "E": 3, "F": 6},
		"E": {"C": 8, "D": 3},
		"F": {"D": 6},
	}

	//fmt.Println("graph", graph)

	distance, parent := dijkstra(graph, "A")

	fmt.Println("parent", parent)
	fmt.Println("distance", distance)

	printLine(distance, parent)

}
