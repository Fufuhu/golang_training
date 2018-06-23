package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)

	ch <- 5
	ch <- 4

	i := <-ch

	fmt.Println(i)

	i = <-ch

	fmt.Println(i)

}
