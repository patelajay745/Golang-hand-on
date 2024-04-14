package main

import (
	"fmt"
	"maps"
)

func main() {

	//map[key-type]val-type
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 113

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//return zero if not found
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	delete(m, "k2")
	fmt.Println("Map after delete:", m)

	clear(m)
	fmt.Println("map:", m)

	val, prs := m["k2"]
	fmt.Println("prs:", prs, val)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n1 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n1) {
		fmt.Println("n==n2")
	}

}
