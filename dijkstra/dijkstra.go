package dijkstra

import (
	"fmt"
	"math"
)

const MAX = math.MaxInt32

var processed = NewSet()

// 狄克斯特拉 算法
func dijkstra(graph map[string]map[string]int, start string) (map[string]int, map[string]string) {

	pq := &PQ{}
	parent := make(map[string]string)
	distance := initDistance(graph, start)
	// fmt.Println("distance is", distance)
	parent[start] = "nil"

	pq.PushKey(start, 0)

	for pq.Len() > 0 {
		v, d := pq.PoPKey()
		processed.Add(v)
		nodes := graph[v]
		// v -> w 的点
		for w, _ := range nodes {
			if !processed.Has(w) {
				if d+graph[v][w] < distance[w] {
					pq.PushKey(w, d+graph[v][w])
					parent[w] = v
					distance[w] = d + graph[v][w]
				}
			}
		}
	}
	return distance, parent
}

func initDistance(graph map[string]map[string]int, start string) map[string]int {
	distance := make(map[string]int)
	distance[start] = 0
	for k, _ := range graph {
		if k != start {
			distance[k] = MAX
		}
	}

	return distance
}

func printLine(distance map[string]int, parent map[string]string) {
	p := parent["F"]
	l := len(parent) - 2

	fmt.Println("F", "--", distance["F"]-distance[p], "->", p)
	for l > 0 {
		fmt.Println(p, "--", distance[p]-distance[parent[p]], "->", parent[p])
		p = parent[p]
		l--
	}
}
