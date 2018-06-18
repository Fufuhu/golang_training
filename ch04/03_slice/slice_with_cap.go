package main

import (
	"fmt"
)

func main() {

	s := make([]int, 5, 10)

	fmt.Println(s)
	s[0] = 1
	s = append(s, 5)
	s = append(s, 6)
	s = append(s, 7)
	s = append(s, 8)
	s = append(s, 9)
	fmt.Println(s)
}
