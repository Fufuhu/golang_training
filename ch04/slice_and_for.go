package main

import (
	"fmt"
)

func main() {
	s := []string{
		"Apple",
		"Banana",
		"Cherry",
	}

	for i, v := range s {
		fmt.Printf("[%d] => %s\n", i, v)
	}
}
