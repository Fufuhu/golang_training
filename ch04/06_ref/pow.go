package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	pow(a)
	fmt.Println(a)
}

func pow(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}
