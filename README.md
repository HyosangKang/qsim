# Quantum Simulator qsim

- Author: Hyosang Kang (hyosang@dgist.ac.kr)
- A personal quantum simulator written by GO

# How to use

- Construct a quantum circuit and show diagram.
```go
c := qsim.NewCircuit(3)
c.CX(0, 2)
c.X(1)
c.Show()
```
- Result
```
StateVector: (1.000+0.000i)|010>

Diagram: [0]-o---
             |
         [1]---X-
             |
         [2]-X---
```