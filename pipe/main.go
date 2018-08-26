package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

func Goodmorning(name string) string {
	return "Good morning, " + name
}

func Bow(message func(string) string) {
	mes := message("hoge")
	fmt.Println(mes)
}

func main() {
	Bow(Hello)
	Bow(Goodmorning)
}
