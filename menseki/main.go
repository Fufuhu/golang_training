package main

import "fmt"

type Size struct {
	width  int
	height int
}

func (s *Size) Triangle() int {
	return s.width * s.height / 2
}

func (s *Size) Rectancle() int {
	return s.width * s.height
}

func main() {

	p := &Size{
		width:  10,
		height: 20,
	}

	fmt.Println(p.Triangle())
	fmt.Println(p.Rectancle())
}
