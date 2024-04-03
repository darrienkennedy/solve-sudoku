package main

import (
	"github.com/darrienkennedy/solve-sudoku/pkg/solver"
)

func main() {
	b := solver.NewFromFile("board00")
	s := solver.NewSolver(b)

	basecase := 4
	for !s.IsCompleted() {
		s.Step()
		s.Print()
		basecase -= 1

		if basecase == 0 {
			break
		}
	}
}
