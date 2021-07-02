package main

import (
	"./circuit"
)

func main() {
	//k1 := braket.NewKet(1, 0)
	//k2 := braket.NewKet(0, 1)
	//b := braket.NewBra(k2)
	//op := braket.NewOp(k1, b)
	//fmt.Println(op.Act(k2))
	c := circuit.NewCircuit(4)
	c.X(1)
	c.CX(3, 0)
	c.X(1)
	c.Show()
}
