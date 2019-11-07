package main

import (
	"image/color"
	"math"

	"fmt"

	"gonum.org/v1/gonum/integrate"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func G(phiI, phiJ, dphiI, dphiJ func(float64) float64) float64 {
	//values := nodes(100)
	g := func(x float64) float64 {
		return k(x)*dphiI(x)*dphiJ(x) + q(x)*phiI(x)*phiJ(x)
	}
	//y := evaluatePoints(g, values)
	//res := integrate.Trapezoidal(values, y)

	m := int(m(g, A, B, 2))
	N := methodRunge(g, A, B, m, 2)
	res := SimpsonsMethod(g, A, B, N)
	res += ALPHA1 * phiI(A) * phiJ(A)
	res += ALPHA2 * phiI(B) * phiJ(B)
	return res
}

func L(v func(float64) float64) float64 {
	//values := nodes(100)
	l := func(x float64) float64 {
		return newRitzF(x) * v(x)
	}
	//y := evaluatePoints(l, values)
	//res := integrate.Trapezoidal(values, y)

	//B := findUpperBound(l, 100)
	m := int(m(l, A, B, 1))
	N := methodRunge(l, A, B, m, 1)
	res := SimpsonsMethod(l, A, B, N)
	return res
}

func Ritz(count int) []float64 {
	E := mat.NewDense(count, count, nil)
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if math.IsNaN(G(phi[i], phi[j], dphi[i], dphi[j])) {
				E.Set(i, j, 0)
			}
			E.Set(i, j, G(phi[i], phi[j], dphi[i], dphi[j]))
		}
	}
	l := make([]float64, count)
	for i := 0; i < count; i++ {
		l[i] = L(phi[i])
	}
	F := mat.NewDense(len(x), 1, l)

	//fmt.Println(E.RawMatrix().Data[0:n])
	//fmt.Println(E.RawMatrix().Data[n : 2*n])
	//fmt.Println(E.RawMatrix().Data[2*n : 3*n])
	fmt.Printf("Ritz cond: %.5f\n", mat.Cond(E, 2))
	var Res mat.Dense
	Res.Solve(E, F)
	res := make([]float64, count)
	for i := range res {
		res[i] = Res.RawRowView(i)[0]
	}

	actual := make([]float64, count)
	c := mat.NewVecDense(count, actual)
	c.MulVec(E, Res.ColView(0))
	fmt.Print(c.RawVector().Data)
	fmt.Print("\n")
	for i := 0; i < count; i++ {
		fmt.Printf("%.7f \t", c.RawVector().Data[i]-l[i])
	}
	fmt.Print("\n")
	return res
}

//func close(value float64) float64 {
//	var answer float64
//	for i, c := range ConstRitz {
//		answer += c * phi[i](value)
//	}
//	return answer + zeroBasic(value)
//}

func showRitz(solution, polinom func(float64) float64) {
	ImageFunc := plotter.NewFunction(solution)
	ImageFunc.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	ImageFunc.Samples = 100

	AprFun := plotter.NewFunction(polinom)
	AprFun.Color = color.RGBA{R: 30, G: 108, B: 153, A: 111}
	AprFun.Samples = 100

	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = A, B
	pl.Y.Min, pl.Y.Max = -2, 1

	pl.Add(ImageFunc)
	pl.Add(AprFun)

	pl.Title.Text = "Approximation"
	pl.Legend.Add("Function", ImageFunc)
	pl.Legend.Add("Aproximation by Ritz", AprFun)
	if err := pl.Save(5*vg.Inch, 5*vg.Inch, "Ritz.png"); err != nil {
		panic(err.Error())
	}
}

func Ganother(phiI, phiJ, dphiI, d2phiI func(float64) float64) float64 {
	values := nodes(100)
	g := func(x float64) float64 {
		return (-1)*k(x)*d2phiI(x)*phiJ(x) + (-1)*dk(x)*dphiI(x)*phiJ(x) + q(x)*phiI(x)*phiJ(x)
	}
	y := evaluatePoints(g, values)
	res := integrate.Trapezoidal(values, y)

	//m := int(m(g, A, B, 2))
	//N := methodRunge(g, A, B, m, 2)
	//res := SimpsonsMethod(g, A, B, N)
	return res
}

func Ritzanother(count int) []float64 {
	E := mat.NewDense(count, count, nil)
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if math.IsNaN(Ganother(phi[i], phi[j], dphi[i], d2phi[i])) {
				E.Set(i, j, 0)
			} else {
				E.Set(i, j, Ganother(phi[i], phi[j], dphi[i], d2phi[i]))
			}
		}
	}
	l := make([]float64, count)
	for i := 0; i < count; i++ {
		l[i] = L(phi[i])
	}
	F := mat.NewDense(len(x), 1, l)

	//fmt.Println(E.RawMatrix().Data[0:n])
	//fmt.Println(E.RawMatrix().Data[n : 2*n])
	//fmt.Println(E.RawMatrix().Data[2*n : 3*n])
	//fmt.Printf("Ritz cond: %.5f\n", mat.Cond(E, 2))
	var Res mat.Dense
	Res.Solve(E, F)
	res := make([]float64, count)
	for i := range res {
		res[i] = Res.RawRowView(i)[0]
	}
	return res
}

//func closeAnother(value float64) float64 {
//	var answer float64
//	for i, c := range ConstRitzAnother {
//		answer += c * phi[i](value)
//	}
//	return answer + zeroBasic(value)
//}

func showRitzAnother(solution, polinom func(float64) float64) {
	ImageFunc := plotter.NewFunction(solution)
	ImageFunc.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	ImageFunc.Samples = 100

	AprFun := plotter.NewFunction(polinom)
	AprFun.Color = color.RGBA{R: 30, G: 108, B: 153, A: 111}
	AprFun.Samples = 100

	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = A, B
	pl.Y.Min, pl.Y.Max = -2, 1

	pl.Add(ImageFunc)
	pl.Add(AprFun)

	pl.Title.Text = "Approximation"
	pl.Legend.Add("Function", ImageFunc)
	pl.Legend.Add("Aproximation by Ritz", AprFun)
	if err := pl.Save(5*vg.Inch, 5*vg.Inch, "RitzAnother.png"); err != nil {
		panic(err.Error())
	}
}
