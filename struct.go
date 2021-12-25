package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, z float64
}

func main() {

	//var c Circle
	cc := Circle{1, 2, 3}
	//ccc := new(Circle)

	fmt.Println(circleArea(&cc))
	fmt.Println(cc.area())
}

func (c *Circle) area() float64 {
	return math.Pi * c.z * c.z
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.z * c.z
}
