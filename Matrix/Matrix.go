package Matrix

import (
	"MatrixGO/Math"
	"fmt"
)

type Martix [][]Math.Fraction

func NewMatrix(rows, columns int) (matrix Martix) {
	matrix = make([][]Math.Fraction, rows)
	for row := 0; row < rows; row++ {
		matrix[row] = make([]Math.Fraction, columns)
		for col := 0; col < columns; col++ {
			matrix[row][col] = Math.NewFracNum(0)
		}
	}
	return
}

func NewScalarMatrix(size, val int) (matrix Martix) {

	matrix = NewMatrix(size, size)

	for rc := 0; rc < size; rc++ {
		matrix[rc][rc] = Math.NewFracNum(val)
	}

	return
}

func MatrixMulMatrix(a, b Martix) (c Martix) {
	c = NewMatrix(a.Rows(), b.Columns())

	for row := 0; row < c.Rows(); row++ {
		for col := 0; col < c.Columns(); col++ {
			for s := 0; s < a.Columns(); s++ {
				frac := Math.FracMulFrac(a[row][s], b[s][col])
				c[row][col] = Math.FracAddFrac(c[row][col], frac).FracCut()
			}
		}
	}

	return
}

func (matrix Martix) Rows() int {
	return len(matrix)
}

func (matrix Martix) Columns() int {
	return len(matrix[0])
}

func (matrix Martix) GetRow(row int) (res []Math.Fraction) {
	res = make([]Math.Fraction, matrix.Columns())

	for col := 0; col < matrix.Columns(); col++ {
		res[col] = matrix[row][col]
	}

	return
}

func (matrix Martix) GetColumn(col int) (res []Math.Fraction) {
	res = make([]Math.Fraction, matrix.Rows())

	for row := 0; row < matrix.Rows(); row++ {
		res[row] = matrix[row][col]
	}

	return
}

func (matrix Martix) Transform() (res Martix) {
	res = NewMatrix(matrix.Columns(), matrix.Rows())

	for row := 0; row < res.Rows(); row++ {
		for col := 0; col < res.Columns(); col++ {
			res[row][col] = matrix[col][row]
		}
	}

	return
}

func (matrix Martix) ToString() (res string) {
	for _, row := range matrix {
		for _, col := range row {
			res += col.ToString() + "\t"
		}
		res += "\n"
	}

	return
}

func (matrix Martix) Print() Martix {
	fmt.Print("\n", matrix.ToString())

	return matrix
}

func (matrix Martix) Println() Martix {
	fmt.Print("\n", matrix.ToString(), "\n")

	return matrix
}
