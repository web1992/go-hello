package maptest

import (
	"fmt"
)

// MapCreate for init map use make
func MapCreate() {

	var myMap map[string]int

	fmt.Println("myMap is", myMap)

	// use make
	myMap = make(map[string]int)
	myMap["k1"] = 1
	fmt.Println("myMap is", myMap)

	// use {}
	myMap2 := map[string]int{}
	myMap2["a"] = 1
	fmt.Println("myMap2 is", myMap2)
	fmt.Println("myMap2[a] is", myMap2["a"])

	myMap3 := map[int]int{1: 1, 2: 2}
	fmt.Println("myMap3 is", myMap3)

}

func MapGet() {

	myMap := map[string]int{"1": 1, "2": 2}

	if v1, ok := myMap["1"]; ok {
		fmt.Println("v1 is", v1)

	}

	if v2, ok := myMap["22"]; ok {
		fmt.Println("v2 is", v2)

	}
	fmt.Println(myMap)
	v22, ok := myMap["22"]
	fmt.Println("key 22 is", v22, "ok is", ok)
}

// MapTest test
func MapTest() {

	map1 := make(map[string]int)
	map1["k1"] = 1

	fmt.Println(map1)
	fmt.Println(map1["k1"])
	fmt.Println(map1["k2"])

	value, exist := map1["kx"]
	fmt.Println(value, exist)

	map2 := make(map[string]string)

	if v, e := map2["k1"]; !e {
		fmt.Println("key k1 not map value =", v)
	}

	map2["k2"] = "v2"
	if v, e := map2["k2"]; e {
		fmt.Println("key k2 in map value =", v)
	}

	map2["k3"] = "v3"
	fmt.Println("key k3 in map value =", map2["k3"])
	delete(map2, "k3")
	fmt.Println("key k3 in map value =", map2["k3"])

	for k, vv := range map2 {
		fmt.Println("k=", k, "v=", vv)
	}

}
