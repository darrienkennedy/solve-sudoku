package solver

import (
	"strconv"
	"testing"
)

// areEqual is a helper function to see if two byte arrays are equal
func areEqual(a, b [9]uint8) bool {
	for i := 0; i < 9; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// bytesToString is a helper function which returns a string from a byte array
func bytesToString(bytes [9]uint8) string {
	var ret string
	for _, b := range bytes {
		ret += strconv.Itoa(int(b))
	}
	return ret
}

// Test_getGroup tests method getGroup of Board
func Test_getGroup(t *testing.T) {
	board := NewFromFile("boardTest")

	makeArrayOf := func(num uint8) [9]uint8 {
		return [9]uint8{
			num, num, num,
			num, num, num,
			num, num, num,
		}
	}

	var d [9]uint8
	for i := 0; i < 9; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			d = board.getGroup(i)
			comp := makeArrayOf(uint8(i + 1))
			if !areEqual(d, comp) {
				t.Errorf("fail")
			}
		})
	}
}

// Test_getCol tests method getCol of Board
func Test_getCol(t *testing.T) {
	board := NewFromFile("boardTest")
	expectCols := [9]string{
		"111444777",
		"111444777",
		"111444777",
		"222555888",
		"222555888",
		"222555888",
		"333666999",
		"333666999",
		"333666999",
	}

	var d [9]uint8
	for i := 0; i < 9; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			d = board.getCol(i)
			if bytesToString(d) != expectCols[i] {
				t.Errorf("fail")
			}
		})
	}
}

// Test_getRow tests method getRow of Board
func Test_getRow(t *testing.T) {
	board := NewFromFile("boardTest")
	expectRows := [9]string{
		"111222333",
		"111222333",
		"111222333",
		"444555666",
		"444555666",
		"444555666",
		"777888999",
		"777888999",
		"777888999",
	}

	var d [9]uint8
	for i := 0; i < 9; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			d = board.getRow(i)
			if bytesToString(d) != expectRows[i] {
				t.Errorf("fail")
			}
		})
	}
}

// TestClone tests if the clone method works as expected
func TestClone(t *testing.T) {
	board := NewFromFile("boardTest")
	clone := board.Clone()

	t.Run("0", func(t *testing.T) {
		eq := board.equals(clone)
		if !eq {
			t.Errorf("fail")
		}
	})

	board.Set(0, 9)
	t.Run("1", func(t *testing.T) {
		eq := board.equals(clone)
		if eq {
			t.Errorf("fail")
		}
	})
}
