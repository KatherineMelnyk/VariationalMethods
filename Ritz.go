package main

import (
	"image/color"
	"math"

	"gonum.org/v1/gonum/integrate"

	"fmt"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func G(phiI, phiJ, dphiI, dphiJ func(float64) float64) float64 {
	values := nodes(100)
	g := func(x float64) float64 {
		return k(x)*dphiI(x)*dphiJ(x) + q(x)*phiI(x)*phiJ(x)
	}
	y := evaluatePoints(g, values)
	res := integrate.Trapezoidal(values, y)
	//res := integrate.Simpsons(values, y)
	//B := findUpperBound(g, 2000)
	//m := int(m(g, A, B, 1))
	//N := methodRunge(g, 0, B, m, 1)
	//res := SimpsonsMethod(g, A, B, N)
	res -= k(3) * dphiJ(3) * phiI(3)
	res += k(1) * dphiJ(1) * phiI(1)
	return res
}

func L(v func(float64) float64) float64 {
	values := nodes(100)
	l := func(x float64) float64 {
		return newRitzF(x) * v(x)
	}
	y := evaluatePoints(l, values)
	res := integrate.Trapezoidal(values, y)
	//res := integrate.Simpsons(values, y)
	//B := findUpperBound(l, 2000)
	//m := int(m(l, A, B, 1))
	//N := methodRunge(l, 0, B, m, 1)
	//SimpsonsMethod(l, A, B, N)
	return res
}

func Ritz(n int) []float64 {
	bas := BasicFunc(n)
	dbas := dBasicFunc(n)
	E := mat.NewDense(n, n, nil)
	//coef := matrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if math.IsNaN(G(bas[i], bas[j], dbas[i], dbas[j])) {
				E.Set(i, j, 0)
			}
			E.Set(i, j, G(bas[i], bas[j], dbas[i], dbas[j]))
		}
	}
	l := make([]float64, n)
	for i := 0; i < n; i++ {
		l[i] = L(bas[i])
	}
	//elem := FromMattoVec(coef)
	F := mat.NewDense(len(x), 1, l)
	//E := mat.NewDense(len(coef), len(coef[0]), elem)

	fmt.Println(E.RawMatrix().Data[0:n])
	fmt.Println(E.RawMatrix().Data[n : 2*n])
	fmt.Println(E.RawMatrix().Data[2*n : 3*n])
	fmt.Printf("cond: %.5f\n", mat.Cond(E, 2))
	var Res mat.Dense
	Res.Solve(E, F)
	res := make([]float64, n)
	for i := 0; i < len(res); i++ {
		res[i] = Res.RawRowView(i)[0]
	}
	return res
}

func close(value float64) float64 {
	var answer float64
	phi := BasicFunc(n)
	for i, c := range ConstRitz {
		answer += c * phi[i](value)
	}
	return answer + zeroBasic(value)
}

func showRitz(solution, polinom func(float64) float64) {
	ImageFunc := plotter.NewFunction(solution)
	ImageFunc.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	//ImageFunc.Width = vg.Inch / 20
	ImageFunc.Samples = 100

	AprFun := plotter.NewFunction(polinom)
	AprFun.Color = color.RGBA{R: 30, G: 108, B: 153, A: 111}
	//AprFun.Width = vg.Inch / 20
	AprFun.Samples = 100

	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = A, B
	pl.Y.Min, pl.Y.Max = -3, 1

	pl.Add(ImageFunc)
	pl.Add(AprFun)

	pl.Title.Text = "Approximation"
	//pl.Title.Font.Size = vg.Inch
	//pl.Legend.Font.Size = vg.Inch / 2
	//pl.Legend.XOffs = -vg.Inch
	//pl.Legend.YOffs = vg.Inch / 2
	pl.Legend.Add("Function", ImageFunc)
	pl.Legend.Add("Aproximation by Collocatiom", AprFun)
	if err := pl.Save(5*vg.Inch, 5*vg.Inch, "Ritz.png"); err != nil {
		panic(err.Error())
	}
}
