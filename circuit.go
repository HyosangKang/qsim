package qsim

import (
	"fmt"
	"math"
	"math/cmplx"
)

type Circuit struct {
	N       int        // number of qubits
	Diagram [][]string // Gates on circuit
	Unitary []Mat
}

func NewCircuit(n int) *Circuit {
	c := Circuit{
		N:       n,
		Unitary: []Mat{},
	}
	c.Diagram = make([][]string, 2*n-1)
	for i := 0; i < n; i++ {
		c.Diagram[i] = []string{}
	}
	return &c
}

func (c *Circuit) X(i int) {
	m := InitOp(c.N)
	m[i] = X()
	c.Unitary = append(c.Unitary, TensorMat(m))
	c.Diagram[2*i] = append(c.Diagram[2*i], "X")
}

func (c *Circuit) H(i int) {
	c.Diagram[i] = append(c.Diagram[i], "H")
}

func (c *Circuit) CX(i, j int) {
	// Add unitary gate
	r := 1
	for k := 0; k < c.N; k++ {
		r *= 2
	}
	id := Id(r)
	m1 := Id(2)
	for k := 0; k < c.N; k++ {
		var n Mat
		if k == i {
			n = Z()
		} else {
			n = I()
		}
		m1 = Tensor(m1, Mat(n))
	}
	m2 := Id(2)
	for k := 0; k < c.N; k++ {
		var n Mat
		if k == j {
			n = X()
		} else {
			n = I()
		}
		m2 = Tensor(m2, Mat(n))
	}
	m3 := Id(2)
	for k := 0; k < c.N; k++ {
		var n Mat
		switch {
		case k == i:
			n = Z()
		case k == j:
			n = X()
		default:
			n = I()
		}
		m3 = Tensor(m3, Mat(n))
	}
	m3 = m3.Phase(math.Pi)

	m := Add(id, m1, m2, m3)
	m = m.Factor(complex(float64(1)/float64(2), 0))
	c.Unitary = append(c.Unitary, m)

	// Add diagram
	if i == j || i >= c.N || j >= c.N {
		panic("Invalid CNOT gate.")
	}
	i1, i2 := i, j
	if i > j {
		i1, i2 = j, i
	}
	n := len(c.Diagram[i1])
	for k := 2 * i1; k <= 2*i2; k += 2 {
		if n < len(c.Diagram[k]) {
			n = len(c.Diagram[k])
		}
	}
	for k := 2 * i1; k <= 2*i2; k += 2 {
		for l := len(c.Diagram[k]); l < n; l++ {
			c.Diagram[k] = append(c.Diagram[k], "-")
		}
	}
	c.Diagram[2*i] = append(c.Diagram[2*i], "o")
	c.Diagram[2*j] = append(c.Diagram[2*j], "X")
	for k := 2*i1 + 2; k <= 2*i2-2; k += 2 {
		c.Diagram[k] = append(c.Diagram[k], "-")
	}
	l := len(c.Diagram[2*i])
	for k := 2*i1 + 1; k <= 2*i2-1; k += 2 {
		for s := len(c.Diagram[k]); s < l-1; s++ {
			c.Diagram[k] = append(c.Diagram[k], " ")
		}
		c.Diagram[k] = append(c.Diagram[k], "|")
	}
}

func (c *Circuit) StateVector() Mat {
	k := TensorMat(InitKet(c.N))
	for _, u := range c.Unitary {
		k = Mul(u, k)
	}
	return k
}

func (c *Circuit) ShowState() {
	m := c.StateVector()
	s := ""
	first := true
	for i := 0; i < len(m); i++ {
		if cmplx.Abs(m[i][0]) < 1e-10 {
			continue
		}
		b := Binary(i, c.N)
		if !first {
			s += "+"
		}
		s += fmt.Sprintf("%.3f", m[i][0])
		s += "|" + b + ">"
		first = false
	}
	fmt.Println(s)
}

func (c *Circuit) Show() {
	l := 0
	for i := 0; i < len(c.Diagram); i++ {
		if len(c.Diagram[i]) > l {
			l = len(c.Diagram[i])
		}
	}
	for i := 0; i < len(c.Diagram); i++ {
		if i%2 == 0 {
			fmt.Printf("[%d]-", i/2)
		} else {
			fmt.Printf("    ")
		}
		for j := 0; j < len(c.Diagram[i]); j++ {
			fmt.Printf("%s", c.Diagram[i][j])
			if i%2 == 0 {
				fmt.Printf("-")
			} else {
				fmt.Printf(" ")
			}
		}
		for j := len(c.Diagram[i]); j < l; j++ {
			if i%2 == 0 {
				fmt.Printf("--")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
}
