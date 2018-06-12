package main

import (
	"fmt"
	"lineargo/matrix"
)

func main() {
	x := matrix.NewMatrix([][]float64{{1, 2}, {2, 1}})
	fmt.Println(x)
}
