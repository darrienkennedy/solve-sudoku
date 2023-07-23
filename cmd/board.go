package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Board struct {
	mu   sync.Mutex
	data [81]uint8
}

// NewFromFile constructs a Board from a file containing board data
func NewFromFile(filename string) *Board {
	bytes, err := os.ReadFile(fmt.Sprintf("../../data/%s", filename))
	if err != nil {
		panic(err)
	}

	var ret Board
	var i uint8
	for _, b := range bytes {
		if b != '\n' {
			ret.data[i] = b - '0'
			i += 1
		}
	}

	return &ret
}

// getRow returns a row from a Board as an array of bytes
func (m *Board) getRow(n int) [9]uint8 {
	if n < 0 || n > 8 {
		panic("getRow, bad input: n")
	}

	var ret [9]uint8
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := 0; i < 9; i++ {
		ret[i] = m.data[9*n+i]
	}

	return ret
}

// getCol returns a column from a Board as an array of bytes
func (m *Board) getCol(n int) [9]uint8 {
	if n < 0 || n > 8 {
		panic("getCol, bad input: n")
	}

	var ret [9]uint8
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := 0; i < 9; i++ {
		ret[i] = m.data[9*i+n]
	}

	return ret
}

// getGroup returns a group from a Board as an array of bytes
func (m *Board) getGroup(n int) [9]uint8 {
	if n < 0 || n > 8 {
		panic("getGroup, bad input: n")
	}

	var ret [9]uint8
	m.mu.Lock()
	defer m.mu.Unlock()
	var i uint8
	var j uint8
	var k uint8 = uint8(n*3 + 18*(n/3))
	for ; i < 9; i++ {
		if i%3 == 0 && i != 0 {
			j += 1
		}
		ret[i] = m.data[uint8(6*j)+k+i]
	}

	return ret
}

// PrintRow prints a Board row in standard formatting
func (m *Board) PrintRow(n int) {
	row := m.getRow(n)
	var outbuf string
	for i, d := range row {
		if i%3 == 0 && i != 0 {
			outbuf += " "
		}
		outbuf += strconv.Itoa(int(d))
	}
	fmt.Println(outbuf)
}

// PrintCol prints a Board column in standard formatting
func (m *Board) PrintCol(n int) {
	col := m.getCol(n)
	var outbuf string
	for i, d := range col {
		if i%3 == 0 && i != 0 {
			outbuf += "\n"
		}
		outbuf += strconv.Itoa(int(d))
		outbuf += "\n"
	}
	fmt.Println(outbuf)
}

// PrintGroup prints a Board group in standard formatting
func (m *Board) PrintGroup(n int) {
	group := m.getGroup(n)
	var outbuf string
	for i, d := range group {
		if i%3 == 0 && i != 0 {
			outbuf += "\n"
		}
		outbuf += strconv.Itoa(int(d))
	}
	fmt.Println(outbuf)
}

// Print prints a Board in standard formatting
func (m *Board) Print() {
	m.mu.Lock()
	defer m.mu.Unlock()

	var outbuf string
	for i, d := range m.data {
		if i == 0 {
			outbuf += "- sudoku --\n"
		} else if i%27 == 0 {
			outbuf += "\n\n"
		} else if i%9 == 0 {
			outbuf += "\n"
		} else if i%3 == 0 {
			outbuf += " "
		}

		outbuf += strconv.Itoa(int(d))
	}
	outbuf += "\n-----------"
	fmt.Println(outbuf)
}
