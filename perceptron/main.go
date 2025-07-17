package main

import "log"

func main() {

	a := []float64{1, 2, 3}
	w := []float64{4, 5, 6}
	b := 2.0
	log.Print(perceptron(a, w, b))
}

func perceptron(x []float64, w []float64, b float64) float64 {
	res := vecCalc(x, w) + b

	return stepFunc(res)
}

func stepFunc(v float64) float64 {
	if v >= 0 {
		return 1
	}
	return 0
}

func vecCalc(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("Vectors must be of the same length")
	}
	var sum float64
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}
