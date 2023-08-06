package main

func main() {
	b := NewFromFile("board00")
	s := NewSolver(b)
	s.Print()
}
