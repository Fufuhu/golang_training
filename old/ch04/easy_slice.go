package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a) //[1, 2, 3, 4, 5]

	s := a[0:2]
	fmt.Println(s) //[1, 2]

	t := a[:4]
	fmt.Println(t) //[1, 2, 3, 4] [0:4]と等価

	u := a[:]
	fmt.Println(u) //[1, 2, 3, 4, 5] [:]は[0:len(u)]と等価
}