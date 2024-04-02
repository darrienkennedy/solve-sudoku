package solver

import (
	"fmt"
	"sync"
)

type Solver struct {
	mu       sync.Mutex
	step     int
	original *Board
	steps    []*Board
	final    *Board
}

// NewSolver creates a solver from a *Board.
func NewSolver(b *Board) *Solver {
	res := Solver{
		original: b.Clone(),
		steps: []*Board{
			b.Clone(),
		},
	}

	return &res
}

// Print prints the solver.
func (m *Solver) Print() {
	fmt.Printf("st: %d\n", m.step)
	m.steps[m.step].Print()
}
