package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func Constructor(_x int, _y int) *Point {
	return &Point{
		x: _x,
		y: _y,
	}
}

func Dist(p1 *Point, p2 *Point) float64 {
	a := float64(p1.x-p2.x)
	b := float64(p1.y-p2.y)
	return math.Sqrt(a*a + b*b)
}

func main(){
	p1 := Constructor(8, 9)
	p2 := Constructor(6, 7)
	fmt.Println(Dist(p1, p2))
}