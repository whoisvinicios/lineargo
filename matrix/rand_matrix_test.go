package matrix

import (
	"fmt"
	"testing"
)

func TestRandMatrix(t *testing.T) {
	m := RandMatrix(10000, 10000)
	fmt.Println(m)
}

func TestRandMatrixConc(t *testing.T) {
	m := RandMatrixConc(10000, 10000)
	fmt.Println(m)
}
