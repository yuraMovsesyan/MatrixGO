package Matrix

import (
	"MatrixGO/STD"
	"fmt"
)

type Matrix [][]STD.Fraction

func NewMatrix(rows, columns int) (matrix Matrix) {
	matrix = make([][]STD.Fraction, rows)
	for row := 0; row < rows; row++ {
		matrix[row] = make([]STD.Fraction, columns)
		for col := 0; col < columns; col++ {
			matrix[row][col] = STD.NewFracNum(0)
		}
	}
	return
}

func NewScalarMatrix(size, val int) (matrix Matrix) {

	matrix = NewMatrix(size, size)

	for rc := 0; rc < size; rc++ {
		matrix[rc][rc] = STD.NewFracNum(val)
	}

	return
}

func MatrixMulMatrix(a, b Matrix) (c Matrix) {
	c = NewMatrix(a.Rows(), b.Columns())

	for row := 0; row < c.Rows(); row++ {
		for col := 0; col < c.Columns(); col++ {
			for s := 0; s < a.Columns(); s++ {
				frac := STD.FracMulFrac(a[row][s], b[s][col])
				c[row][col] = STD.FracAddFrac(c[row][col], frac).FracCut()
			}
		}
	}

	return
}

func (matrix Matrix) Rows() int {
	return len(matrix)
}

func (matrix Matrix) Columns() int {
	return len(matrix[0])
}

func (matrix Matrix) GetRow(row int) (res []STD.Fraction) {
	res = make([]STD.Fraction, matrix.Columns())

	for col := 0; col < matrix.Columns(); col++ {
		res[col] = matrix[row][col]
	}

	return
}

func (matrix Matrix) GetColumn(col int) (res []STD.Fraction) {
	res = make([]STD.Fraction, matrix.Rows())

	for row := 0; row < matrix.Rows(); row++ {
		res[row] = matrix[row][col]
	}

	return
}

func (matrix Matrix) Transform() (res Matrix) {
	res = NewMatrix(matrix.Columns(), matrix.Rows())

	for row := 0; row < res.Rows(); row++ {
		for col := 0; col < res.Columns(); col++ {
			res[row][col] = matrix[col][row]
		}
	}

	return
}

func (matrix Matrix) ToString() (res string) {
	for _, row := range matrix {
		for _, col := range row {
			res += col.ToString() + "\t"
		}
		res += "\n"
	}

	return
}

func (matrix Matrix) Print() Matrix {
	fmt.Print("\n", matrix.ToString())

	return matrix
}

func (matrix Matrix) Println() Matrix {
	fmt.Print("\n", matrix.ToString(), "\n")

	return matrix
}
