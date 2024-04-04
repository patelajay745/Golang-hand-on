package main

import "fmt"

func main() {
	fmt.Println("welcome to example of my array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[3] = "banana"

	var vegList = [5]string{"tomato", "potato"}

	fmt.Println("fruit list : ", fruitList)
	fmt.Println("veggie list : ", vegList)

}
