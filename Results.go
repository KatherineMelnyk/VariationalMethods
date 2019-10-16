package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const A = 1.
const B = 3.
const ALPHA1, ALPHA2 = 1., 1.
const Q1, Q2, Q3 = 3., 2., 1.
const K1, K2, K3 = 3., 3., 2.
const P1, P2, P3 = 1., 1., 1.
const M1, M2, M3 = 1., 2., 0.

const n = 11

var x = nodes(n)

func main() {
	//printVector(x)
	printMatrix(systemConstants(n, x))
	ImageFunc := plotter.NewFunction(u)
	ImageFunc.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	ImageFunc.Width = vg.Inch / 20
	ImageFunc.Samples = 100

	AprFun := plotter.NewFunction(polinom)
	AprFun.Color = color.RGBA{R: 30, G: 108, B: 153, A: 111}
	AprFun.Width = vg.Inch / 20
	AprFun.Samples = 100

	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = A, B
	pl.Y.Min, pl.Y.Max = -1, 1

	pl.Add(ImageFunc)
	pl.Add(AprFun)

	pl.Title.Text = "Approximation"
	pl.Title.Font.Size = vg.Inch
	pl.Legend.Font.Size = vg.Inch / 2
	pl.Legend.XOffs = -vg.Inch
	pl.Legend.YOffs = vg.Inch / 2
	pl.Legend.Add("Function", ImageFunc)
	pl.Legend.Add("Aproximation by Collocatiom", AprFun)
	if err := pl.Save(14*vg.Inch, 14*vg.Inch, "Task2.png"); err != nil {
		panic(err.Error())
	}

}
