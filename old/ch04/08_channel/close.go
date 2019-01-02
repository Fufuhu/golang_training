package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	var (
		i  int
		ok bool
	)
	ok = true

	for ok {
		i, ok = <-ch
		fmt.Print(i)
		fmt.Println(ok)
	}
}
