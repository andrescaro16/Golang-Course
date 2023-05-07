package figures

import "fmt"

type figures2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

func NewSquare(base float64) square {
	return square{ base: base }
}

func NewRectangle(base, height float64) rectangle {
	return rectangle{ base: base, height: height}
}

func (s square) area() float64 {
	return s.base * s.base
}

func (r rectangle) area() float64 {
	return r.base * r.height
}

func Calculate(f figures2D) {
	fmt.Println("Area:", f.area())
}
