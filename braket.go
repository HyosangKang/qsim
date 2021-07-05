package qsim

import (
	"math"
)

func NewKet(i int) Mat {
	m := NewMat(2, 1)
	switch {
	case i == 0:
		m[0][0] = 1
	case i == 1:
		m[1][0] = 1
	default:
		panic("Invalid Ket")
	}
	return m
}

func NewBra(i int) Mat {
	m := NewMat(1, 2)
	switch {
	case i == 0:
		m[0][0] = 1
	case i == 1:
		m[0][1] = 1
	default:
		panic("Invalid Bra")
	}
	return m
}

func NewOp(k, b Mat) Mat {
	r := len(k[0])
	c := len(b)
	if r != c {
		panic("Invalid Operation from Ket-Bra")
	}
	return Mul(k, b)
}

func X() Mat {
	m := NewOp(NewKet(0), NewBra(1))
	n := NewOp(NewKet(1), NewBra(0))
	return Add(m, n)
}

func I() Mat {
	m := NewOp(NewKet(0), NewBra(0))
	n := NewOp(NewKet(1), NewBra(1))
	return Add(m, n)
}

func Z() Mat {
	m := NewOp(NewKet(0), NewBra(0))
	n := NewOp(NewKet(1).Phase(math.Pi), NewBra(1))
	return Add(m, n)
}

func H() Mat {
	k0, k1 := NewKet(0), NewKet(1)
	m := NewOp(Add(k0, k1), NewBra(0))
	n := NewOp(Add(k0, k1.Phase(math.Pi)), NewBra(1))
	return Add(m, n)
}
