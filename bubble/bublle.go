package main

import "fmt"

func sort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := len(array) - 1; j > i; j-- {
			if compare(array[j-1], array[j]) {
				temp := array[j-1]
				array[j-1] = array[j]
				array[j] = temp
			}
		}
	}
}

func compare(a int, b int) bool {
	if a > b {
		return true
	}
	return false
}

func main() {
	array := []int{2, 4, 3, 5, 1}
	sort(array)
	for i, v := range array {
		fmt.Printf("%d, %d\n", i, v)
	}
}
