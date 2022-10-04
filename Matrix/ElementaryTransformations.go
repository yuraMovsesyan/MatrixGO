package Matrix

import "MatrixGO/STD"

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

func (matrix Matrix) RowMulFrac(row int, frac STD.Fraction) Matrix {

	for colum := 0; colum < matrix.Columns(); colum++ {
		matrix[row][colum] = STD.FracMulFrac(matrix[row][colum], frac).FracCut()
	}

	return matrix
}

func (matrix Matrix) ColumMulFrac(colum int, frac STD.Fraction) Matrix {

	for row := 0; row < matrix.Rows(); row++ {
		matrix[row][colum] = STD.FracMulFrac(matrix[row][colum], frac).FracCut()
	}

	return matrix
}

func (matrix Matrix) RowAddRowMulFrac(row1, row2 int, fracForRow2 STD.Fraction) Matrix {

	for colum := 0; colum < matrix.Columns(); colum++ {
		matrix[row1][colum] = STD.FracAddFrac(matrix[row1][colum], STD.FracMulFrac(matrix[row2][colum], fracForRow2)).FracCut()
	}

	return matrix
}

func (matrix Matrix) ColumAddColumMulFrac(colum1, colum2 int, fracForColum2 STD.Fraction) Matrix {

	for row := 0; row < matrix.Rows(); row++ {
		matrix[row][colum1] = STD.FracAddFrac(matrix[row][colum1], STD.FracMulFrac(matrix[row][colum2], fracForColum2)).FracCut()
	}

	return matrix
}
