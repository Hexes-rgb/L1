package main

import (
	"fmt"
	"math"
)

type Point struct { //define Point structure
	x, y float64
}

func NewPoint(x, y float64) Point { //Point constructor
	return Point{
		x: x,
		y: y,
	}
}

func (point1 Point) DistanceTo(point2 Point) float64 { //method for get distance between two points
	dx := point2.x - point1.x
	dy := point2.y - point1.y
	return math.Sqrt(dx*dx + dy*dy) //mathematical formula for calculating the distance between two points
}

func main() {
	point1 := NewPoint(1, 1)
	point2 := NewPoint(2, 2)
	fmt.Println(point1.DistanceTo(point2))
}
