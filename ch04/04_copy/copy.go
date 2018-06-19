package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 11}

	n := copy(s1, s2)

	fmt.Println(n)
	fmt.Println(s1)

}
