package circuit

import "fmt"

type Circuit struct {
	N    int        // number of qubits
	Gate [][]string // Gates on circuit
}

func NewCircuit(n int) *Circuit {
	return &Circuit{
		N:    n,
		Gate: [][]string{},
	}
}

func (c *Circuit) X(i int) {
	c.Gate[i] = append(c.Gate[i], "X")
}

func (c *Circuit) H(i int) {
	c.Gate[i] = append(c.Gate[i], "X")
}

func (c *Circuit) CX(i, j int) {
	n, m := len(c.Gate[i]), len(c.Gate[j])
	if n < m {
		for i := 0; i < m-n; i++ {
			c.Gate[i] = append(c.Gate[i], "-")
		}
	} else {
		for i := 0; i < m-n; i++ {
			c.Gate[j] = append(c.Gate[j], "-")
		}
	}
	c.Gate[i] = append(c.Gate[i], "o")
	c.Gate[j] = append(c.Gate[j], "X")
}

func (c *Circuit) Show() {
	l := 0
	for i := 0; i < len(c.Gate); i++ {
		if len(c.Gate[i]) > l {
			l = len(c.Gate[i])
		}
	}
	for i := 0; i < len(c.Gate); i++ {
		for j := 0; j < len(c.Gate[i]); j++ {
			fmt.Printf("%s", c.Gate[i][j])
		}
	}
}
