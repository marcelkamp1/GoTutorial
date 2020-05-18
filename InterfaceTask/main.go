package main

import "fmt"

type triangle struct {
	height float64
	base float64
}
type square struct{
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main()  {
	t := triangle{}
	s := square{}

	printArea(t)
	printArea(s)
}

func (triangle) getArea() float64{
	height := 4.0
	base := 3.0
	areaOfTriangle := 0.5 * base * height
	return areaOfTriangle
}

func (square) getArea() float64 {
	sideLength := 2.0
	result := sideLength * sideLength
	return result
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
