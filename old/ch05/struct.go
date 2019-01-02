package main

import "fmt"

func main() {
	type Feed struct {
		Name   string
		Amount int
	}

	type Animal struct {
		Name string
		Feed Feed
	}

	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	fmt.Println(a.Name)
	fmt.Println(a.Feed.Name)
	fmt.Println(a.Feed.Amount)
}
