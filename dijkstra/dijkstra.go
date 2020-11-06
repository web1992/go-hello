package dijkstra

import "fmt"

const MAX = ^int32(0)

// 狄克斯特拉 算法
func dijkstra() {

	m := make(map[string]map[string]int32)

	sm := map[string]int32{"A": 6, "B": 2}
	am := map[string]int32{"end": 1}
	bm := map[string]int32{"A": 3, "end": 5}
	em := map[string]int32{}

	m["start"] = sm
	m["A"] = am
	m["B"] = bm
	m["end"] = em

	costMap := make(map[string]int32)
	costMap["A"] = 6
	costMap["B"] = 2
	costMap["end"] = MAX

	parentMap := make(map[string]string)
	parentMap["A"] = "start"
	parentMap["B"] = "start"
	parentMap["end"] = ""

	fmt.Println(m)
	fmt.Println(costMap)

	minNode := find_min_cost(costMap)

	fmt.Println("minNode", minNode)

	for minNode != "" {

		c := costMap[minNode]

		nM := m[minNode]

		keys := getKeys(nM)

		for _, k := range keys {
			new_cost := c + nM[k]

			if costMap[k] > new_cost {
				costMap[k] = new_cost
				parentMap[k] = minNode
			}
		}

		minNode = find_min_cost(costMap)
	}

	fmt.Println(parentMap["end"])
	fmt.Println(parentMap[parentMap["end"]])
	fmt.Println(parentMap[parentMap[parentMap["end"]]])
	fmt.Println(parentMap[parentMap[parentMap[parentMap["end"]]]])

}

func getKeys(m map[string]int32) []string {

	var arr []string

	for k, _ := range m {
		arr = append(arr, k)
	}

	return arr
}

func find_min_cost(cost map[string]int32) string {

	min := MAX
	n := ""

	for k, v := range cost {
		if v < min {
			min = v
			n = k
		}
	}

	return n
}
