package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("Array a :", a)

	a[4] = 100
	fmt.Println("Setting up value:", a)
	fmt.Println("To get a value:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array B:", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j

		}
	}

	fmt.Println("2d Array:", twoD)

}
