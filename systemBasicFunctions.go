package main

import "math"

func C() float64 {
	gamma := k(B)
	return B + ((gamma * (B - A)) / (ALPHA2*(B-A) + 2*gamma))
}

func D() float64 {
	beta := k(A)
	return A - ((beta * (B - A)) / (ALPHA1*(B-A) + 2*beta))
}

func Apsi() float64 {
	return (mu1(A) - ALPHA1*mu2(B)/ALPHA2) /
		(-k(A) + ALPHA1*A - (ALPHA1/ALPHA2)*(k(B)+ALPHA2*B))
}

func Bpsi() float64 {
	return (mu2(B) - Apsi()*(k(B)+ALPHA2*B)) / ALPHA2
}

func zeroBasic(x float64) float64 {
	return Apsi()*x + Bpsi()
}

func dzeroBasic(x float64) float64 {
	return Apsi()
}

func d2zeroBasic(x float64) float64 {
	return 0.
}

func firstBasic(x float64) float64 {
	return math.Pow(x-A, 2) * (x - C())
}

func dfirstBasic(x float64) float64 {
	return 2.*(x-A)*(x-C()) + math.Pow(x-A, 2)
}

func d2firstBasic(x float64) float64 {
	return 2*(x-C()) + 4*(x-A)
}

func secondBasic(x float64) float64 {
	return math.Pow(B-x, 2) * (x - D())
}

func dsecondBasic(x float64) float64 {
	return -2.*(B-x)*(x-D()) + math.Pow(B-x, 2)
}

func d2secondBasic(x float64) float64 {
	return 2*(x-D()) - 4*(B-x)
}

func system(x float64, i int) float64 {
	return math.Pow(x-A, float64(i-1)) * math.Pow(B-x, 2)
}

func BasicFunc(n int) []func(float64) float64 {
	var phi []func(float64) float64
	phi = append(phi, firstBasic)
	phi = append(phi, secondBasic)
	for i := 2; i < n; i++ {
		i := i
		phi = append(phi, func(x float64) float64 {
			return system(x, i)
		})
	}
	return phi
}

func dsystem(x float64, i int) float64 {
	return float64(i-1)*math.Pow(x-A, float64(i-2))*math.Pow(B-x, 2) + math.Pow(x-A, float64(i-1))*(B-x)*float64(-2)

}

func dBasicFunc(n int) []func(float64) float64 {
	var phi []func(float64) float64
	phi = append(phi, dfirstBasic)
	phi = append(phi, dsecondBasic)
	for i := 2; i < n; i++ {
		i := i
		phi = append(phi, func(x float64) float64 {
			return dsystem(x, i)
		})
	}
	return phi
}

func d2system(x float64, i int) float64 {
	return float64(i-1)*float64(i-2)*math.Pow(x-A, float64(i-3))*math.Pow(B-x, 2) + float64(i-1)*math.Pow(x-A, float64(i-2))*(B-x)*float64(-2) + float64(i-1)*math.Pow(x-A, float64(i-2))*float64(-2)*(B-x) + float64(2)*math.Pow(x-A, float64(i-1))

}

func d2BasicFunc(n int) []func(float64) float64 {
	var phi []func(float64) float64
	phi = append(phi, d2firstBasic)
	phi = append(phi, d2secondBasic)
	for i := 2; i < n; i++ {
		i := i
		phi = append(phi, func(x float64) float64 {
			return d2system(x, i)
		})
	}
	return phi
}
