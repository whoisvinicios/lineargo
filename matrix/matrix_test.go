package matrix

import (
	"testing"
)

func TestMatrix_New(t *testing.T) {
	x := NewMatrix([][]float64{{2, 1}, {1, 2}})
	if x == nil {
		t.Fatal("NewMatrix are returnuing nil matrix")
	}
}

func TestMatrix_Create(t *testing.T) {
	x := CreateMatrix(5, 5)
	if x == nil {
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