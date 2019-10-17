package main

import (
	"image/color"
	"math"

	"fmt"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func systemConstants(n int, x []float64) [][]float64 {
	sysConst := matrix(n, len(x))
	phi := BasicFunc(n)
	dphi := dBasicFunc(n)
	d2phi := d2BasicFunc(n)
	for i := range x {
		for j := range phi {
			sysConst[i][j] -= dk(x[i]) * dphi[j](x[i])
			sysConst[i][j] -= k(x[i]) * d2phi[j](x[i])
			sysConst[i][j] += p(x[i]) * dphi[j](x[i])
			sysConst[i][j] += q(x[i]) * phi[j](x[i])
			if math.IsNaN(sysConst[i][j]) {
				sysConst[i][j] = 0
			}
		}
	}
	return sysConst
}

func Collocation(n int, x []float64) []float64 {
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
	fmt.Print(E.RawMatrix().Data)
	res := make([]float64, len(valueFunc))
	for i := 0; i < len(res); i++ {
		res[i] = Res.RawRowView(i)[0]
	}
	return res
}

func polinom(value float64) float64 {
	var answer float64
	phi := BasicFunc(n)
	for i, c := range ConstCollocation {
		answer += c * phi[i](value)
	}
	return answer + zeroBasic(value)
}

func showCollocation(solution, polinom func(float64) float64) {
	ImageFunc := plotter.NewFunction(solution)
	ImageFunc.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	ImageFunc.Width = vg.Inch / 19
	ImageFunc.Samples = 100

	AprFun := plotter.NewFunction(polinom)
	AprFun.Color = color.RGBA{R: 30, G: 108, B: 153, A: 111}
	AprFun.Width = vg.Inch / 19
	AprFun.Samples = 100

	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = A, B
	pl.Y.Min, pl.Y.Max = -1, 1

	pl.Add(ImageFunc)
	pl.Add(AprFun)

	pl.Title.Text = "Approximation"
	pl.Title.Font.Size = vg.Inch / 1.5
	pl.Legend.Font.Size = vg.Inch / 2.5
	pl.Legend.XOffs = -vg.Inch
	pl.Legend.YOffs = vg.Inch / 3
	pl.Legend.Add("Function", ImageFunc)
	pl.Legend.Add("Aproximation by Collocatiom", AprFun)
	if err := pl.Save(10*vg.Inch, 10*vg.Inch, "Collocation.png"); err != nil {
		panic(err.Error())
	}
}
