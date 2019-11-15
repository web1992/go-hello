package maptest

import (
	"fmt"
)

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
