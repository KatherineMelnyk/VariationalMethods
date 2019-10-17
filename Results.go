package main

const A = 1.
const B = 3.
const ALPHA1, ALPHA2 = 1., 1.
const Q1, Q2, Q3 = 3., 2., 1.
const K1, K2, K3 = 3., 3., 2.
const P1, P2, P3 = 1., 1., 1.
const M1, M2, M3 = 1., 2., 0.

const n = 30

var x = nodes(n)
var ConstRitz = Ritz(n)
var ConstCollocation = Collocation(n, x)

func main() {
	showCollocation(u, polinom)
	showRitz(u, close)
}
