package matrix

import "testing"

func TestCreateMatrix(t *testing.T) {
	x := NewMatrix([][]float64{{2, 1}, {1, 2}})

	if x.matrix == nil {
		t.Fatal("CreateMatrix are returnuing nil matrix")
	}
}

func TestMatrix_Compare(t *testing.T) {
	x := NewMatrix([][]float64{{2, 1}, {1, 2}})
	y := NewMatrix([][]float64{{2, 1}, {1, 2}})

	if !x.Compare(y) {
		t.Fatal("Error")
	}
}