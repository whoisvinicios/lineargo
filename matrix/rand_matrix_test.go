package matrix

import (
	"testing"
)

var x = RandMatrix(1000, 1000)
var y = RandMatrix(1000, 1000)

func TestRandMatrix(t *testing.T) {
	x.Multiply(y)
}

func TestRandMatrixConc(t *testing.T) {
	x.MultiplyConcurrent(y)
}
