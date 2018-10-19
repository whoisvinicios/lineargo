package matrix

// Add two matrices
func (m *Matrix) Add(o *Matrix) *Matrix {
	if (m.rows == o.rows) && (m.columns == o.columns) {
		panic("The dimensions of the matrices are not the same")
	}
	out := make([][]float64, m.rows)
	for x := 0; x < m.rows; x++ {
		out[x] = make([]float64, o.columns)
		for y := 0; y < o.columns; y++ {
			out[x][y] += m.matrix[x][y] + o.matrix[x][y]
		}
	}
	return &Matrix{matrix: out, rows: m.rows, columns: o.columns}
}
