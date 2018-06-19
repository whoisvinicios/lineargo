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

func TestVector_ZerozVector(t *testing.T) {
	v := ZeroVector(3)
	if !v.Compare(NewVector([]float64{0, 0, 0})) {
		t.Fatal("The vectors are not equals")
	}
}
