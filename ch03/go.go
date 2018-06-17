package main

import (
	"fmt"
)

func main() {
	go sub()
	go func() {
		for j := 0; ; j++ {
			fmt.Printf("j=%d\n", j)
		}
	}()
	for {
		fmt.Println("main")
	}
}

func sub() {
	for i := 0; ; i++ {
		fmt.Printf("i=%d\n", i)
	}
}
