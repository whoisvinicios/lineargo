package matrix

import (
	"testing"
	"fmt"
)

func TestMatrixMultiplication(t *testing.T) {
	a := Matrix{}
	x := NewMatrix([][]float64{{1, 2}, {2, 1}})
	//y := Matrix{matrix: [][]float64{{4, 2}, {6, 1}}, rows: 2, columns: 2}
	//result := Matrix{matrix: [][]float64{{16, 4}, {14, 5}}, rows: 2, columns: 2}
	//
	fmt.Println(a)
	fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(result)
	//
	//z := x.Multiply(y)
	//
	//eq := result.Compare(z)
	//
	//if !eq{
	//	t.Fatal("Not equals")
	//}
}
