package main

import (
	"MatrixGO/Matrix"
	"MatrixGO/STD"
	"fmt"
)

func main() {
	a := Matrix.Matrix{
		{STD.NewFracNum(4)},
		{STD.NewFracNum(-1)},
		{STD.NewFracNum(-2)},
	}

	id, _ := STD.MaxFracArray(a.GetColumn(0))

	fmt.Println(id)
}
