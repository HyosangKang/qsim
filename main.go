package main

import (
	"./braket"
	"fmt"
)

func main() {
	k1 := braket.NewKet(1, 0)
	k2 := braket.NewKet(0, 1)
	b := braket.NewBra(k2)
	op := braket.NewOp(k1, b)
	fmt.Println(op.Act(k2))
}
