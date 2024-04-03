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

// Step performs one iteration of checks of spaces which can be filled out.
func (m *Solver) Step() {
	return
}

// IsCompleted checks if the board has any squares with values of 0.
func (m *Solver) IsCompleted() bool {
	cb := m.steps[m.step]
	for _, v := range cb.data {
		if v == 0 {
			return false
		}
	}

	return true
}

// Print prints the solver.
func (m *Solver) Print() {
	fmt.Printf("st: %d\n", m.step)
	m.steps[m.step].Print()
}
