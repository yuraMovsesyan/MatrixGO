package main

import (
	"MatrixGO/Matrix"
)

func main() {
	a := Matrix.NewScalarMatrix(3, 2)
	b := Matrix.NewScalarMatrix(3, 3)

	Matrix.MatrixMulMatrix(a, b).Println()
}
