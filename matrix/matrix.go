package matrix

// An Matrix represents value distributed in rows and column
type Matrix struct {
	matrix  [][]float64
	rows    int
	columns int
}

// NewMatrix returns a matrix of given slice of slices
// NewMatrix([][]float64{{1, 2}, {2, 1}})
func NewMatrix(m [][]float64) *Matrix {
	return &Matrix{matrix: m, rows: len(m), columns: len(m[0])}
}

// CreateMatrix return a new matrix of zeros
func CreateMatrix(row, columns int) *Matrix {
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
func (m Matrix) Multiply(o Matrix) *Matrix {
	if m.columns != o.rows {
		panic("The number of rows must be equals to the number of rows of the matrix")
	}
	out := make([][]float64, len(m.matrix))
	for x := range m.matrix {
		for y := range o.matrix {
			out[x] = make([]float64, len(o.matrix))
			out[x][y] += m.matrix[x][y] * o.matrix[y][x]
		}
	}
	return &Matrix{matrix: out, rows: m.rows, columns: m.columns}
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
