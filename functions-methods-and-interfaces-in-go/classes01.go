package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p Point) distanceToOrigin() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

func (p *Point) OffsetX(offset float64) {
	p.x += offset
}

func main() {
	x1 := Point{5, 2}
	fmt.Println(x1.distanceToOrigin())
	x1.OffsetX(2)
	fmt.Println(x1)
}
