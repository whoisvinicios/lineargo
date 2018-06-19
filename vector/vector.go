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

// ZeroVector return a new 1D array of zeros
func ZeroVector(n int) *Vector {
	return &Vector{vector: make([]float64, n), length: n}
}

// Add two vector
func (v *Vector) Add(o *Vector) *Vector {
	out := make([]float64, v.length)
	for x := range v.vector {
		out[x] = v.vector[x] + o.vector[x]
	}
	return &Vector{vector: out, length: len(out)}
}

// ScalarMultiply a
func (v *Vector) ScalarMultiply(o float64) *Vector {
	out := make([]float64, v.length)
	for x := range v.vector {
		out[x] = o * v.vector[x]
	}
	return &Vector{vector: out, length: len(out)}
}

// Compare the given vector
func (v *Vector) Compare(o *Vector) bool {
	for x := range v.vector {
		if v.vector[x] != o.vector[x] {
			return false
		}
	}
	return true
}
