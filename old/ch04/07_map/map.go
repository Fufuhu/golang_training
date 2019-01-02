package main

import "fmt"

func main() {
	m := make(map[int]string)

	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"

	fmt.Println(m)

	n := make(map[string]string)

	n["Yamada"] = "Taro"
	n["Sato"] = "Hanako"
	n["Yamada"] = "Jiro"

	fmt.Println(n)
}
