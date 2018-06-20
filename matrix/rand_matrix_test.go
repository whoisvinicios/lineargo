package matrix

import (
	"testing"
)

var x = RandMatrix(10000, 10000)
var y = RandMatrix(10000, 10000)

func TestRandMatrix(t *testing.T) {
	x.Multiply(y)
}

func TestRandMatrixConc(t *testing.T) {
	x.MultiplyConcurrent(y)
}
