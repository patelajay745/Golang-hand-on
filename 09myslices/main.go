package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices example")

	var fruitList = []string{"apple", "mongo"}

	fruitList = append(fruitList, "Banana", "cherry")
	fmt.Println("Fruitlist:", fruitList)

	fruitList = append(fruitList[1:])

	fmt.Println("Fruitlist:", fruitList)

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores = append(highscores, 955, 666)

	sort.Ints(highscores)

	fmt.Println("Highsocres: ", highscores)

	courses := []string{"reactjs", "python", "java", "kotlin"}
	index := 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
