package main

import (
	"fmt"
	"math"
)

// in this version I am using math, look at version one to see it doen manually
// ---but an interesting way to construct a const value for a struct in go

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	a:= s.side * s.side
	return a
}

func (c circle) area()float64 {
	a := c.radius * math.Pi
	return a
}

type shape interface {
	area() float64
}

func shapeInfo(s shape){
	switch s.(type){
	case circle:
		fmt.Println("This shape is a circle and has an area of:", s.area(), "and a radius of:", s.(circle).radius)

	case square:
		fmt.Println("This shappe is a square with an area of", s.area(), "and a side length of:", s.(square).side)
	// fmt.Println("This shape has an area of", s.area())
	}
}

func main(){
	sq1 := square {
		side: 5,
	}
	sq1Area:= sq1.area()
	fmt.Println(sq1Area)

	cir1 := circle {
		radius: 10,
	}
	cir1Area:= cir1.area()
	fmt.Println(cir1Area)

	shapeInfo(sq1)
	shapeInfo(cir1)
}