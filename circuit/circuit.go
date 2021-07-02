package circuit

import (
	"../braket"
	"fmt"
)

type Circuit struct {
	N       int        // number of qubits
	Diagram [][]string // Gates on circuit
	Unitary []braket.Mat
}

func NewCircuit(n int) *Circuit {
	c := Circuit{
		N:       n,
		Unitary: []braket.Mat{},
	}
	c.Diagram = make([][]string, 2*n-1)
	for i := 0; i < n; i++ {
		c.Diagram[i] = []string{}
	}
	return &c
}

func (c *Circuit) X(i int) {
	m := braket.Identity(2)
	for j := 0; j < c.N; j++ {
		var n braket.Mat
		if j == i {
			n = braket.X()
		} else {
			n = braket.Identity(2)
		}
		m = braket.Tensor(n, m)
	}
	c.Unitary = append(c.Unitary, m)
	c.Diagram[2*i] = append(c.Diagram[2*i], "X")
}

func (c *Circuit) H(i int) {
	c.Diagram[i] = append(c.Diagram[i], "H")
}

func (c *Circuit) CX(i, j int) {
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
