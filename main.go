package main

import (
	matrix "github.com/LFLATH/GoMath/Matrix"
)

func main() {

	a := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	b := [][]float64{
		{1, 2, 3},
		{0, 0, 0},
		{3, 4, 5},
	}
	matrix.Display(a)
	c, err := matrix.Dot(a, b)
	if err != nil {
		panic(err)
	}
	matrix.Display(c)
	d := matrix.Transpose(c)
	matrix.Display(d)

	e := matrix.Sum(d)
}
