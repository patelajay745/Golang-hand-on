package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("unintilized ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appended", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copied slice:", c)

	l := s[2:5]
	fmt.Println("slice1:", l)

	l1 := s[2:]
	fmt.Println("slice2", l1)

	t := []string{"g", "h", "i"}
	fmt.Println("declared:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t==t2")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)

}
