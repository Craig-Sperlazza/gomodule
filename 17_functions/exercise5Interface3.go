package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	rad float64
}

func (s square)area() float64 {
	return s.length * s.length
}

func (c circle)area() float64 {
	return ((c.rad * c.rad) * math.Pi)
}

type shape interface {
	area() float64
}

func shapeInfo(s shape){
	switch s.(type){
	case circle:
		fmt.Println("Your circle has a radius of", s.(circle).rad, "and an area of", s.area())

	case square:
		fmt.Println("Your square has a side length of", s.(square).length, "and an area of", s.area())
	}
}

func main(){
	fmt.Println(2 * math.Pi)

	s1 := square {
		length: 5,
	}

	c1 := circle{
		rad: 2,
	}

	fmt.Println(s1.area())

	fmt.Println(c1.area())

	shapeInfo(s1)
	shapeInfo(c1)
}
