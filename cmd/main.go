package main

import (
	"github.com/darrienkennedy/solve-sudoku/pkg/solver"
)

func main() {
	b := solver.NewFromFile("board00")
	s := solver.NewSolver(b)
	s.Print()
}
