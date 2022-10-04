package Matrix

import "MatrixGO/Math"

func (matrix Matrix) SwapRows(row1, row2 int) Matrix {

	for colum := 0; colum < matrix.Columns(); colum++ {
		matrix[row1][colum], matrix[row2][colum] = matrix[row2][colum], matrix[row1][colum]
	}

	return matrix
}

func (matrix Matrix) SwapColumns(colum1, colum2 int) Matrix {

	for row := 0; row < matrix.Rows(); row++ {
		matrix[row][colum1], matrix[row][colum2] = matrix[row][colum2], matrix[row][colum1]
	}

	return matrix
}

func (matrix Matrix) RowMulFrac(row int, frac Math.Fraction) Matrix {

	for colum := 0; colum < matrix.Columns(); colum++ {
		matrix[row][colum] = Math.FracMulFrac(matrix[row][colum], frac).FracCut()
	}

	return matrix
}

func (matrix Matrix) ColumMulFrac(colum int, frac Math.Fraction) Matrix {

	for row := 0; row < matrix.Rows(); row++ {
		matrix[row][colum] = Math.FracMulFrac(matrix[row][colum], frac).FracCut()
	}

	return matrix
}

func (matrix Matrix) RowAddRowMulFrac(row1, row2 int, fracForRow2 Math.Fraction) Matrix {

	for colum := 0; colum < matrix.Columns(); colum++ {
		matrix[row1][colum] = Math.FracAddFrac(matrix[row1][colum], Math.FracMulFrac(matrix[row2][colum], fracForRow2)).FracCut()
	}

	return matrix
}

func (matrix Matrix) ColumAddColumMulFrac(colum1, colum2 int, fracForColum2 Math.Fraction) Matrix {

	for row := 0; row < matrix.Rows(); row++ {
		matrix[row][colum1] = Math.FracAddFrac(matrix[row][colum1], Math.FracMulFrac(matrix[row][colum2], fracForColum2)).FracCut()
	}

	return matrix
}
