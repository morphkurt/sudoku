package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `.3.562.41
..48.....
6...497..
.86...32.
...6.7...
.75...19.
..921...8
.....65..
56.738.1.`

func TestParser(t *testing.T) {

	got := Parse(i)
	want := [][]int{[]int{0, 3, 0, 5, 6, 2, 0, 4, 1}, []int{0, 0, 4, 8, 0, 0, 0, 0, 0}, []int{6, 0, 0, 0, 4, 9, 7, 0, 0}, []int{0, 8, 6, 0, 0, 0, 3, 2, 0}, []int{0, 0, 0, 6, 0, 7, 0, 0, 0}, []int{0, 7, 5, 0, 0, 0, 1, 9, 0}, []int{0, 0, 9, 2, 1, 0, 0, 0, 8}, []int{0, 0, 0, 0, 0, 6, 5, 0, 0}, []int{5, 6, 0, 7, 3, 8, 0, 1, 0}}

	assert.Equal(t, want, got)
}

func TestReadFile(t *testing.T) {

	got := ReadFile("test_data")
	want := i

	assert.Equal(t, want, got)

}

func TestBoardPrint(t *testing.T) {

	board := Parse(i)

	got := Print(board)
	want := "+-------+-------+-------+\n| . 3 . | 5 6 2 | . 4 1 |\n| . . 4 | 8 . . | . . . |\n| 6 . . | . 4 9 | 7 . . |\n+-------+-------+-------+\n| . 8 6 | . . . | 3 2 . |\n| . . . | 6 . 7 | . . . |\n| . 7 5 | . . . | 1 9 . |\n+-------+-------+-------+\n| . . 9 | 2 1 . | . . 8 |\n| . . . | . . 6 | 5 . . |\n| 5 6 . | 7 3 8 | . 1 . |\n+-------+-------+-------+"

	assert.Equal(t, want, got)

}
