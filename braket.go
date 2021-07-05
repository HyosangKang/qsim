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

func InitKet(n int) []Mat {
	ks := make([]Mat, n)
	for i := 0; i < n; i++ {
		ks[i] = NewKet(0)
	}
	return ks
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

func InitOp(n int) []Mat {
	op := make([]Mat, n)
	for i := 0; i < n; i++ {
		op[i] = I()
	}
	return op
}

func X() Mat {
	m := NewMat(2, 2)
	m[0][1], m[1][0] = 1, 1
	return m
}

func I() Mat {
	m := NewMat(2, 2)
	m[0][0], m[1][1] = 1, 1
	return m
}

func Z() Mat {
	m := I()
	m[1][1] = -1
	return m
}

func H() Mat {
	m := NewMat(2, 2)
	m[0][0], m[1][0] = complex(1/math.Sqrt(2), 0), complex(1/math.Sqrt(2), 0)
	m[0][1], m[1][1] = complex(1/math.Sqrt(2), 0), complex(-1/math.Sqrt(2), 0)
	return m
}
