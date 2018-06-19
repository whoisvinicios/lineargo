package matrix

import (
	"fmt"
	"testing"
)

func TestMatrix_New(t *testing.T) {
	x := NewMatrix([][]float64{{2, 1}, {1, 2}})
	if x == nil {
		t.Fatal("NewMatrix are returnuing nil matrix")
	}
}

func TestMatrix_Multiplication(t *testing.T) {
	x := NewMatrix([][]float64{{2, 3}, {1, 5}})
	y := NewMatrix([][]float64{{4, 3, 6}, {1, -2, 3}})
	c := NewMatrix([][]float64{{11, 0, 21}, {-1, 13, -9}})
	z := x.Multiply(y)
	fmt.Println(z)
	if !z.Compare(c) {
		t.Fatal("The z are not equals to c")
	}
}

func TestMatrix_Create(t *testing.T) {
	x := ZeroMatrix(5, 5)
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
