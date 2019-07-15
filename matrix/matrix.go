package matrix

// Tensor: Multidimensional Matrix
// Struct tensor or Matrix???

// An Matrix represents 2D array
type Matrix struct {
	matrix  [][]float64
	rows    int
	columns int
}

// NewMatrix returns a 2d array of given slice of slices
// NewMatrix([][]float64{{1, 2}, {2, 1}})
func NewMatrix(matrix [][]float64) *Matrix {
	return &Matrix{
		matrix:  matrix,
		rows:    len(matrix),
		columns: len(matrix[0]),
	}
}

// ZeroMatrix return a new matrix of zeros
func ZeroMatrix(row, columns int) *Matrix {
	matrix := make([][]float64, row)
	for i := range matrix {
		matrix[i] = make([]float64, columns)
	}
	return NewMatrix(matrix)
}

// Identity returns an array with only 1 on its main diagonal
func Identity(row, column int) *Matrix {
	return &Matrix{}
}
