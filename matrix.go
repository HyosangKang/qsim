package qsim

import (
	"math"
	"math/cmplx"
)

type Mat [][]complex128

func NewMat(r, c int) Mat {
	m := make([][]complex128, r)
	for i := 0; i < r; i++ {
		m[i] = make([]complex128, c)
	}
	return m
}

func Id(n int) Mat {
	m := make([][]complex128, n)
	for i := 0; i < n; i++ {
		m[i] = make([]complex128, n)
		m[i][i] = 1
	}
	return m
}

func (m Mat) Duplicate() Mat {
	n := make([][]complex128, len(m))
	for i := 0; i < len(m); i++ {
		n[i] = make([]complex128, len(m[0]))
		for j := 0; j < len(m[0]); j++ {
			n[i][j] = m[i][j]
		}
	}
	return n
}

func (m Mat) Phase(t float64) Mat {
	n := m.Duplicate()
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n[0]); j++ {
			n[i][j] *= cmplx.Rect(1, t)
		}
	}
	return n
}

func Add(ms ...Mat) Mat {
	m := ms[0].Duplicate()
	for i := 1; i < len(ms); i++ {
		for j := 0; j < len(m); j++ {
			for k := 0; k < len(m[0]); k++ {
				m[j][k] += ms[i][j][k]
			}
		}
	}
	for j := 0; j < len(m); j++ {
		for k := 0; k < len(m[0]); k++ {
			m[j][k] /= complex(math.Sqrt(float64(len(ms))), 0)
		}
	}
	return m
}

func Tensor(m1, m2 Mat) Mat {
	r1, c1 := len(m1), len(m1[0])
	r2, c2 := len(m2), len(m2[0])
	m := make([][]complex128, r1*r2)
	for i1 := 0; i1 < r1; i1++ {
		for i2 := 0; i2 < r2; i2++ {
			i := i1*r2 + i2
			m[i] = make([]complex128, c1*c2)
			for j1 := 0; j1 < c1; j1++ {
				for j2 := 0; j2 < c2; j2++ {
					j := j1*c2 + j2
					m[i][j] = m1[i1][j1] * m2[i2][j2]
				}
			}
		}
	}
	return m
}

func Mul(m1, m2 Mat) Mat {
	r1, c1 := len(m1), len(m1[0])
	r2, c2 := len(m2), len(m2[0])
	if c1 != r2 {
		panic("Matrix size does not match.")
	}
	m := make([][]complex128, r1)
	for i := 0; i < r1; i++ {
		m[i] = make([]complex128, c2)
		for j := 0; j < c2; j++ {
			a := complex(0, 0)
			for k := 0; k < c1; k++ {
				a += m1[i][k] * m2[k][j]
			}
			m[i][j] = a
		}
	}
	return m
}
