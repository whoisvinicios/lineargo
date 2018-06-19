package matrix

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

// Add two arrays
// TODO implement
func (m Matrix) Add(o Matrix) *Matrix {
	return nil
}

// Subtract two arrays
// TODO implement
func (m Matrix) Subtract(o Matrix) *Matrix {
	return nil
}

// Multiply two arrays
func (m *Matrix) Multiply(o *Matrix) *Matrix {
	if m.columns != o.rows {
		panic("The number of rows must be equals to the number of rows of the matrix")
	}
	out := ZeroMatrix(m.rows, o.columns)
	for x := 0; x < m.rows; x++ {
		for y := 0; y < o.columns; y++ {
			for z := 0; z < o.rows; z++ {
				out.matrix[x][y] += m.matrix[x][z] * o.matrix[z][y]
			}
		}
	}
	return out
}

// ScalarMultiply an array by an scalar number
// TODO inplement
func (m Matrix) ScalarMultiply(e int) *Matrix {
	return nil
}

// CompareConcurrently two matrix concurrently
// for each row this method create a goroutine
// TODO Check if this implementation is the best
func (m *Matrix) CompareConcurrently(o *Matrix) bool {
	for x := range m.matrix {
		go func(x int) bool {
			for y := range o.matrix[0] {
				if m.matrix[x][y] != o.matrix[x][y] {
					return false
				}
			}
			return true
		}(x)
	}
	return true
}

// Compare two matrix
func (m *Matrix) Compare(o *Matrix) bool {
	for x := range m.matrix {
		for y := range o.matrix[0] {
			if m.matrix[x][y] != o.matrix[x][y] {
				return false
			}
		}
	}
	return true
}
