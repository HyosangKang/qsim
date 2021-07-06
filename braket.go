package qsim

import "github.com/hyosangkang/qsim/matrix"

type Ket [][]complex128

func NewKet(n int) Ket {
	m := make([][]complex128, n)
	for i := 0; i < n; i++ {
		m[i] = make([]complex128, 2)
		m[i][0] = 1
	}
	return m
}

func (k Ket) Matrix() matrix.Mat {
	if len(k) == 0 {
		panic("Invalid ket")
	}
	m := matrix.NewMat(2, 1)
	m[0][0], m[1][0] = k[0][0], k[0][1]
	for i := 1; i < len(k); i++ {
		n := matrix.NewMat(2, 1)
		n[0][0], n[1][0] = k[0][0], k[0][1]
		m = matrix.Tensor(n, m)
	}
	return m
}

type Op []matrix.Mat

func NewOp(n int) Op {
	op := make([]matrix.Mat, n)
	for i := 0; i < n; i++ {
		op[i] = matrix.Id(2)
	}
	return op
}

func (o Op) Matrix() matrix.Mat {
	if len(o) == 0 {
		panic("Invalid operation")
	}
	m := o[0]
	for i := 1; i < len(o); i++ {
		m = matrix.Tensor(o[i], m)
	}
	return m
}
