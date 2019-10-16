package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func systemConstants(n int, x []float64) [][]float64 {
	sysConst := matrix(n, len(x))
	phi := BasicFunc(n)
	dphi := dBasicFunc(n)
	d2phi := d2BasicFunc(n)
	i := 0
	for i < len(x) {
		for j := 0; j < len(phi); j++ {
			sysConst[i][j] -= dk(x[i]) * dphi[j](x[i], j)
			sysConst[i][j] -= k(x[i]) * d2phi[j](x[i], j)
			sysConst[i][j] += p(x[i]) * dphi[j](x[i], j)
			sysConst[i][j] += q(x[i]) * phi[j](x[i], j)
			if sysConst[i][j] == math.NaN() {
				sysConst[i][j] = 0
			}
		}
		i++
	}
	return sysConst
}

func Colocation(n int, x []float64) []float64 {
	elements := systemConstants(n, x)
	valueFunc := make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		valueFunc[i] = newF(x[i])
	}
	elem := FromMattoVec(elements)
	F := mat.NewDense(len(x), 1, valueFunc)
	E := mat.NewDense(len(elements), len(elements[0]), elem)
	var Res mat.Dense
	Res.Solve(E, F)
	res := make([]float64, len(valueFunc))
	for i := 0; i < len(res); i++ {
		res[i] = Res.RawRowView(i)[0]
	}
	return res
}

func polinom(value float64) float64 {
	var answer float64
	phi := BasicFunc(n)
	C := Colocation(n, x)
	for i := 0; i < len(C); i++ {
		answer += C[i] * phi[i](value, i)
	}
	return answer + zeroBasic(value)
}
