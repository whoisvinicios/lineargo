package matrix

// This function multiply two arrays
// TODO Try to make concurrent multiplication
func (m *Matrix) Multiply(o *Matrix) *Matrix {
	if m.columns != o.rows {
		panic("The number of rows must be equals to the number of rows of the matrix")
	}
	out := make([][]float64, m.rows)
	for x := 0; x < m.rows; x++ {
		out[x] = make([]float64, o.columns)
		for y := 0; y < o.columns; y++ {
			for z := 0; z < o.rows; z++ {
				out[x][y] += m.matrix[x][z] * o.matrix[z][y]
			}
		}
	}
	return &Matrix{matrix: out, rows: m.rows, columns: o.columns}
}
