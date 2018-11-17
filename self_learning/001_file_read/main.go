package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	buf, err := ioutil.ReadFile("./file.yaml")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", buf)
}
