package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go receiver(ch)
	go receiver2(ch)

	i := 0
	for i < 10000 {
		ch <- i
		i++
	}
}

// int型のチャンネルを受け取る
func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func receiver2(ch <-chan int) {
	for {
		i := <-ch
		fmt.Print("Test:")
		fmt.Println(i)
	}
}
