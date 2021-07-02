package braket

type Mat [][]complex128

func Identity(n int) Mat {
	m := make([][]complex128, n)
	for i := 0; i < n; i++ {
		m[i] = make([]complex128, n)
		m[i][i] = 1
	}
	return m
}

func X() Mat {
	m := make([][]complex128, 2)
	for i := 0; i < 2; i++ {
		m[i] = make([]complex128, 2)
	}
	m[0][1], m[1][0] = 1, 1
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
