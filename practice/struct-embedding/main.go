package main

import "fmt"

type computerBranch struct {
	studentId int
}

func (ce computerBranch) describe() string {
	return fmt.Sprintf("base with num=%v", ce.studentId)
}

type college struct {
	computerBranch
	collegeName string
}

func main() {

	co := college{
		computerBranch: computerBranch{
			studentId: 1,
		},
		collegeName: "McGill",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.studentId, co.collegeName)

	fmt.Println("also num:", co.computerBranch.studentId)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

}
