package main

import (
	"fmt"
)

func main() {
	if x, y := 1, 2; x < y {
		fmt.Printf("%d is less than %d.\n", x, y)
	}

	x, y := 3, 5
	if n := x * y; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	} else {
		fmt.Printf("%d is odd\n", n)
	}

	for i, r := range "ABC" {
		fmt.Printf("[%d] -> %d\n", i, r)
	}

	for i, r := range "あいうえお" {
		fmt.Printf("[%d] -> %d\n", i, r)
	}
}
