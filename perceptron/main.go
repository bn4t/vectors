package main

import "log"

var data = []struct {
	input  []float64
	output float64
}{
	{[]float64{1, 1, 1}, 1},
	{[]float64{1, 1, 0}, 1},
	{[]float64{1, 0, 0}, 1},
	{[]float64{0, 0, 0}, 0},
}

var learningRate = 0.1

func main() {

	weights := []float64{1, 1, 1}
	bias := 0.0

	for _, v := range data {
		yhat := perceptron(v.input, weights, bias)
		epsilon := v.output - yhat

		switch epsilon {
		case 0:
			continue
		case 1:
			// TODO
		case -1:
			// TODO
		}
	}

	a := []float64{1, 2, 3}
	w := []float64{4, 5, 6}
	b := 2.0
	log.Print(perceptron(a, w, b))
}

func perceptron(x []float64, w []float64, b float64) float64 {
	res := dot(x, w) + b

	return stepFunc(res)
}

func stepFunc(v float64) float64 {
	if v >= 0 {
		return 1
	}
	return 0
}

func dot(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("Vectors must be of the same length")
	}
	var sum float64
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}
