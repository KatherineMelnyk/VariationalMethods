package main

import (
	"fmt"
	"math"
)

func evaluatePoints(g func(float64) float64, x []float64) []float64 {
	y := make([]float64, len(x))
	for i := range x {
		y[i] = g(x[i])
	}
	return y
}

func matrix(c, r int) [][]float64 {
	var matrix [][]float64
	for i := 0; i < r; i++ {
		matrix = append(matrix, make([]float64, c))
	}
	return matrix
}

func FromMattoVec(matrix [][]float64) []float64 {
	Size := len(matrix) * len(matrix[0])
	vector := make([]float64, Size)
	for i := 0; i < len(matrix); i++ {
		copy(vector[i*len(matrix[i]):(i+1)*len(matrix[i])], matrix[i])
	}
	return vector
}

func nodes(n int) []float64 {
	nodes := make([]float64, n)
	h := math.Abs(B-A) / float64(n)
	nodes[0] = A
	for i := 1; i < n; i++ {
		nodes[i] = nodes[i-1] + h
	}
	nodes[n-1] = B
	return nodes
}

func printValues(v []func(float64, int) float64, x float64) {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%.8f \t", v[i](x, i))
	}
}

func printMatrix(m [][]float64) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			fmt.Printf("%.4f \t", m[i][j])
		}
		fmt.Print("\n")
	}
}

func printVector(v []float64) {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%.4f \t", v[i])
	}
	fmt.Print("\n")
}
