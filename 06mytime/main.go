package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()

	fmt.Println(presentTime.Format("01/02/2006 15:04:05 Monday"))

	createdDate := time.Date(2025, time.December, 20, 15, 30, 00, 0, time.Local)
	fmt.Println(createdDate.Format("01/02/2006 15:04:05 Monday"))
}
