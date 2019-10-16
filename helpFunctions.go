package main

import (
	"fmt"
	"math"
)

func matrix(c, r int) [][]float64 {
	var matrix [][]float64
	for i := 0; i < r; i++ {
		matrix = append(matrix, []float64{})
		for j := 0; j < c; j++ {
			matrix[i] = append(matrix[i], 0.)
		}
	}
	return matrix
}

func FromMattoVec(matrix [][]float64) []float64 {
	Size := len(matrix) * len(matrix[0])
	vector := make([]float64, Size)
	t := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			vector[t] = matrix[i][j]
			t = t + 1
		}
	}
	return vector
}

func nodes(n int) []float64 {
	nodes := make([]float64, n)
	h := math.Abs(B-A) / float64(n)
	nodes[0] = A + h/2
	for i := 1; i < n; i++ {
		nodes[i] = nodes[i-1] + h
	}
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
