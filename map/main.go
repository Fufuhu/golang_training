package main

import "fmt"

type URL struct {
	urlhandler map[string]func(int, int) int
}

func (url *URL) set(u string, handler func(int, int) int) {

	if url.urlhandler == nil {
		url.urlhandler = make(map[string]func(int, int) int)
	}

	url.urlhandler[u] = handler
}

func (url *URL) get(u string) func(int, int) int {

	if url.urlhandler == nil {
		return nil
	}

	return url.urlhandler[u]
}

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

func main() {
	url := new(URL)

	url.set("plus", plus)
	url.set("minus", minus)
	url.set("multiply", multiply)
	url.set("divide", divide)

	//fmt.Println(url.urlhandler["plus"](1, 2))
	f := url.get("plus")
	fmt.Println(f(1, 2))
}
