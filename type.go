package main

import "fmt"

type point struct {
	x, y int
}

func NewPoint(x, y int) point {
	return point{x, y}
}

type rect struct {
	pos    point
	width  int
	height int
}

func (r rect) Area() int {
	return r.width * r.height
}

func main() {

	p := point{20, 40}
	fmt.Println(p)

	r := rect{
		pos:    NewPoint(20, 40),
		width:  100,
		height: 200,
	}
	fmt.Println(r)

	fmt.Println(r.Area())
}
