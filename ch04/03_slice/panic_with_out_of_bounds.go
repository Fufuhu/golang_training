package main

import (
	"fmt"
)

func main() {
	a := make([]float64, 3)

	fmt.Println(a)
	a[0] = 3.14
	fmt.Println(a)
	a[1] = 6.28
	fmt.Println(a)

	fmt.Println(a[0])
	fmt.Println(a[3]) //panic
}
