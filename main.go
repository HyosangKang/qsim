package main

import (
	"./qsim"
	"fmt"
)

func main() {
	c := qsim.NewCircuit(3)
	c.CX(0, 2)
	c.X(1)
	fmt.Println(c.Unitary)
	c.Show()
}
