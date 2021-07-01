package braket

type Ket []complex128

func NewKet(a, b complex128) Ket {
	return []complex128{a, b}
}

func (k Ket) Mat() Mat {
	m := make([][]complex128, 2)
	m[0] = make([]complex128, 1)
	m[1] = make([]complex128, 1)
	m[0][0], m[1][0] = k[0], k[1]
	return m
}

type Bra []complex128

func NewBra(k Ket) Bra {
	b := make([]complex128, 2)
	b[0], b[1] = k[0], k[1]
	return b
}

type Bras []Bra

type Kets []Ket

func (ks Kets) Mat() Mat {
	if len(ks) < 1 {
		panic("Not enough kets.")
	}
	m := ks[0].Mat()
	for i := 1; i < len(ks); i++ {
		m = Tensor(ks[i].Mat(), m)
	}
	return m
}

type Op [][]complex128

func NewOp(k Ket, b Bra) Op {
	op := make([][]complex128, 2)
	for i := 0; i < 2; i++ {
		op[i] = make([]complex128, 2)
		for j := 0; j < 2; j++ {
			op[i][j] = k[i] * b[j]
		}
	}
	return op
}

func (op Op) Act(k Ket) Ket {
	m := Mul(Mat(op), k.Mat())
	return NewKet(m[0][0], m[1][0])
}

type Ops []Op

func (ops Ops) Mat() Mat {
	if len(ops) < 1 {
		panic("Not enough number of operators")
	}
	m := Mat(ops[0])
	for i := 1; i < len(ops); i++ {
		m = Tensor(Mat(ops[i]), m)
	}
	return m
}
