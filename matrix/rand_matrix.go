package matrix

import (
	"math/rand"
)

func RandMatrix(r, c int) *Matrix {
	rnd := rand.New(rand.NewSource(99))
	out := make([][]float64, r)
	for x := 0; x < r; x++ {
		out[x] = make([]float64, c)
		for y := 0; y < c; y++ {
			out[x][y] = rnd.Float64()
		}
	}
	return &Matrix{
		matrix:  out,
		rows:    r,
		columns: c}
}

func RandMatrixConc(r, c int) *Matrix {
	rnd := rand.New(rand.NewSource(99))
	var out [][]float64
	if r > 100 && c > 100 {
		out := make([][]float64, r)
		for x := 0; x < r; x++ {
			out[x] = make([]float64, c)
			for y := 0; y < c; y++ {
				out[x][y] = rnd.Float64()
			}
		}
	} else {
		out := make([][]float64, r)
		for x := 0; x < r; x++ {
			go func(x int) {
				out[x] = make([]float64, c)
				for y := 0; y < c; y++ {
					out[x][y] = rnd.Float64()
				}
			}(x)
		}
	}
	return &Matrix{
		matrix:  out,
		rows:    r,
		columns: c}
}
