package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Data struct {
	Users []User `yaml:"users"`
}

type User struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Age         int    `yaml:"age"`
}

func main() {
	buf, err := ioutil.ReadFile("./file.yaml")

	if err != nil {
		panic(err)
	}

	var d Data
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		panic(err)
	}
	fmt.Printf("d: %v\n", d)

	for _, user := range d.Users {
		fmt.Printf("%v\n", user)
	}

}
