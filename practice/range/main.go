package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	for _, value := range nums {
		sum += value
	}
	fmt.Println("Sum:", sum)

	for i, nums := range nums {
		if nums == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "Cherry"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("keys:", k)
	}

	// /Unicode
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
