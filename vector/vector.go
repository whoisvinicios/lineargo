package vector

// Vector structure represents 1D array
type Vector struct {
	vector []float64
	length int
}

// NewVector returns a 1D array from the given slice
// NewVector([]float64{1, 2, 3})
func NewVector(v []float64) *Vector {
	return &Vector{vector: v, length: len(v)}
}
