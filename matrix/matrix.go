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
func NewMatrix(m [][]float64) *Matrix {
	return &Matrix{matrix: m, rows: len(m), columns: len(m[0])}
}

// ZeroMatrix return a new matrix of zeros
func ZeroMatrix(row, columns int) *Matrix {
	matrix := make([][]float64, row)
	for i := range matrix {
		matrix[i] = make([]float64, columns)
	}
	return &Matrix{matrix: matrix, rows: row, columns: columns}
}

// Identity returns an array with only 1 on its main diagonal
func Identity(row, column int) *Matrix {
	return &Matrix{}
}
