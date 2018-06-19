package vector

import (
	"testing"
)

func TestVector_NewVector(t *testing.T) {
	v := NewVector([]float64{1, 2, 3})
	if v == nil {
		t.Fatal("The v vector are nil")
	}
}
