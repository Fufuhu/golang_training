package main

import "fmt"

func main() {
	n := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(n)

	//可変長配列にスライスを突っ込むことができる
	s := []int{1, 2, 3}
	n = sum(s...)
	fmt.Println(n)
}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
