package main

import "fmt"

type Shape interface {
	Sides() int
	Area() int
}
type Square struct {
	len int
}

func (s *Square) Sides() int {
	return 4
}

type Square2 struct {
	len int
}

func (s2 *Square2) Sides() int {
	return s2.len
}

func (s2 *Square2) Area() int {
	return s2.len * s2.len
}

func Shaper(s Shape) {
	fmt.Printf("I have sides", s.Sides(), s.Area())
}

func main() {
	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
	// Shaper(&s) // cannot use s (variable of type Square) as Shape value in argument to Shaper: Square does not implement Shape (missing method Area)

	s2 := Square2{len: 4}
	Shaper(&s2)

}
