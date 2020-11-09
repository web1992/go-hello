package dijkstra

import (
	"testing"
)

func Test_dijkstra(t *testing.T) {
	//dijkstra()
}

func Test(t *testing.T) {

	graph := map[string]map[string]int{
		"A": {"B": 5, "C": 1},
		"B": {"D": 1, "C": 2},
		"C": {"B": 2, "D": 4, "E": 8},
		"D": {"B": 1, "C": 4, "E": 3, "F": 6},
		"E": {"C": 8, "D": 3},
		"F": {"D": 6},
	}

	//fmt.Println("graph", graph)

	dijkstra(graph, "A")
}
