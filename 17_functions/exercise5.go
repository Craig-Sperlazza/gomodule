package main

import "fmt"

// this functiomn returns pi which can then be called in the circle area function
// there is no way to put a const in a go struct so this is a good option
func pi() float64 {
	return 3.14159
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	width float64
}

func (s square)area() float64 {
	a := s.length * s.length
	return a
}

func (c circle)area() float64 {
	a :=  pi() * (c.radius * c.radius)
	return a
}

func (r rectangle)area() float64{
	a := r.length * r.width
	return a
}

// interface ----use the interface in the info function below
type shape interface {
	area() float64
}

func info(s shape){
	fmt.Println(s.area())
}

func main(){
	fmt.Println("--------Exercise 5: square, circle---interface area-------------------")
	c1 := circle {
		radius: 4.3,
	}
	circArea1 := c1.area()
	fmt.Println(circArea1)

	sq1 := square {
		length: 10,
	}

	sqArea1 := sq1.area()
	fmt.Println(sqArea1)

	rec1 := rectangle{
		length: 10,
		width: 5,
	}

	recArea1 := rec1.area()
	fmt.Println(recArea1)

	info(c1)
	info(sq1)
	info(rec1)

}