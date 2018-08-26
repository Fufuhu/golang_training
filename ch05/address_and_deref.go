package main

import (
	"fmt"
)

func inc(p *int) {
	*p++
}

func main() {
	i := 1
	inc(&i)
	inc(&i)
	inc(&i)
	fmt.Println(i)
}
