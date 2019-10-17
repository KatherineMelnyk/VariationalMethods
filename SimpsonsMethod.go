package main

import "math"

const EPS = 0.00001

//func F2(x float64) float64 {
//	return math.Pow(math.E, d*x) * math.Sin(c*x)
//}

func points(a, b float64, n int) []float64 {
	x := make([]float64, 2*n+1)
	h := (b - a) / (2 * float64(n))
	for i := 0; i <= 2*n; i++ {
		x[i] = a + float64(i)*h
	}
	return x
}

func quadraticFormula(Func func(x float64) float64, x1, x2, x3 float64) float64 {
	return ((x3 - x1) / 6) * (Func(x1) + 4*Func(x2) + Func(x3))
}

func SimpsonsMethod(Func func(float64) float64, a, b float64, n int) float64 {
	X := points(a, b, n)
	res := 0.
	for i := 0; i < 2*n; i = i + 2 {
		res += quadraticFormula(Func, X[i], X[i+1], X[i+2])
	}
	return res
}

func findUpperBound(Func func(float64) float64, n int) float64 {
	H := 10000.
	B := A + 1
	for math.Abs(SimpsonsMethod(Func, B, H, n)) >= EPS/2 {
		B++
	}
	return B
}

func methodRunge(g func(float64) float64, a, b float64, m, n int) int {
	Ih := SimpsonsMethod(g, a, b, n)
	Ih2 := SimpsonsMethod(g, a, b, 2*n)
	R0h2 := (Ih2 - Ih) / (math.Pow(2, float64(m)) - 1)
	if math.Abs(R0h2) > EPS {
		return methodRunge(g, a, b, m, 2*n)

	}
	return n
}

func m(Func func(float64) float64, a, b float64, n int) float64 {
	Ih := SimpsonsMethod(Func, a, b, n)
	Ih2 := SimpsonsMethod(Func, a, b, 2*n)
	Ih4 := SimpsonsMethod(Func, a, b, 4*n)
	m := (Ih2 - Ih) / (Ih4 - Ih2)
	return math.Log2(m)
}
