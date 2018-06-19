package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := a[2:4]
	fmt.Println(len(s1)) //2(a[2], a[3])
	fmt.Println(cap(s1)) //10 - 2 => 8
	fmt.Println(s1)      //[3,4]

	s2 := a[2:4:4]
	fmt.Println(len(s2)) //2(a[2], a[3])
	fmt.Println(cap(s2)) //4 - 2 => 2
	fmt.Println(s2)      //[3,4]

	s3 := a[2:4:6]
	fmt.Println(len(s3)) //2(a[2], a[3])
	fmt.Println(cap(s3)) //6 - 2 => 4
	fmt.Println(s3)      //[3,4]

}
