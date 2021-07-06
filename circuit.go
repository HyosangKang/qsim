package qsim

import (
	"fmt"
	"github.com/hyosangkang/qsim/matrix"
	"github.com/hyosangkang/qsim/util"
	"math"
	"math/cmplx"
)

type Circuit struct {
	N       int        // number of qubits
	Diagram [][]string // Gates on circuit
	State   matrix.Mat
}

func NewCircuit(n int) *Circuit {
	c := Circuit{
		N:     n,
		State: NewKet(n).Matrix(),
	}
	c.Diagram = make([][]string, 2*n-1)
	for i := 0; i < n; i++ {
		c.Diagram[i] = []string{}
	}
	return &c
}

func (c *Circuit) X(i int) {
	o := NewOp(c.N)
	o[i] = matrix.X()
	c.State = matrix.Mul(o.Matrix(), c.State)
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
	o := matrix.Id(r)

	m := NewOp(c.N)
	m[i] = matrix.Z()
	o = matrix.Add(o, m.Matrix())

	m = NewOp(c.N)
	m[j] = matrix.X()
	o = matrix.Add(o, m.Matrix())

	m = NewOp(c.N)
	m[i] = matrix.Z()
	m[j] = matrix.X()
	o = matrix.Add(o, m.Matrix().Phase(math.Pi))

	c.State = matrix.Mul(o.Factor(complex(.5, 0)), c.State)

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

func (c *Circuit) Show() {
	m := c.State
	s := ""
	first := true
	for i := 0; i < len(m); i++ {
		if cmplx.Abs(m[i][0]) < 1e-10 {
			continue
		}
		b := util.Binary(i, c.N)
		if !first {
			s += "+"
		}
		s += fmt.Sprintf("%.3f", m[i][0])
		s += "|" + b + ">"
		first = false
	}
	fmt.Println("StateVector: " + s + "\n")

	fmt.Printf("Diagram: ")
	l := 0
	for i := 0; i < len(c.Diagram); i++ {
		if len(c.Diagram[i]) > l {
			l = len(c.Diagram[i])
		}
	}
	for i := 0; i < len(c.Diagram); i++ {
		if i != 0 {
			fmt.Printf("         ")
		}
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
