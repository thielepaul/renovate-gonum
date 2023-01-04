package main

import (
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func main() {
	zero := mat.NewDense(3, 5, nil)
	data := make([]float64, 36)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	a := mat.NewDense(6, 6, data)
	zero.Copy(a)
}
