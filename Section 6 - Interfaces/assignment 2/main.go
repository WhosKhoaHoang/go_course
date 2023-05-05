package main

import "fmt"

// Interfaces
type shape interface {
	getArea() float64
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}
// Structs
type square struct{
	sideLength float64
}
type triangle struct{
	base float64
	height float64
}
// Receiver functions
func (s square) getArea() float64 {
	return s.sideLength * 2
}
func (t triangle) getArea() float64 {
	return (0.5)*t.base*t.height 
}

func main() {
	s := square{ sideLength: 2 }
	t := triangle{ base: 2, height: 2 }
	printArea(s)
	printArea(t)
}