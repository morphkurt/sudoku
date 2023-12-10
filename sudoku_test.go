package main

import (
	"testing"

	"github.com/morphkurt/sudoku/util"
	"github.com/stretchr/testify/assert"
)

func TestIsSolved(t *testing.T) {

	unsolvedBoard := `.3.562.41
..48.....
6...497..
.86...32.
...6.7...
.75...19.
..921...8
.....65..
56.738.1.`

	solvedBoard := `756389124
132476985
489521367
371952846
598164273
624837519
915648732
243715698
867293451`

	b := util.Parse(solvedBoard)
	gotSolved := IsSolved(b)
	ub := util.Parse(unsolvedBoard)
	gotUnsolved := IsSolved(ub)

	assert.Equal(t, true, gotSolved)
	assert.Equal(t, false, gotUnsolved)

}

func TestPossibleValues(t *testing.T) {
	bs := `.3.562.41
..48.....
6...497..
.86...32.
...6.7...
.75...19.
..921...8
.....65..
56.738.1.`
	b := util.Parse(bs)
	v0 := GetPossibleValues(b, 0, 0)
	v1 := GetPossibleValues(b, 2, 0)
	v2 := GetPossibleValues(b, 3, 3)
	v3 := GetPossibleValues(b, 6, 6)
	v4 := GetPossibleValues(b, 8, 8)
	v5 := GetPossibleValues(b, 0, 4)
	v6 := GetPossibleValues(b, 4, 4)
	v7 := GetPossibleValues(b, 6, 0)

	assert.Equal(t, []int{7, 8, 9}, v0)
	assert.Equal(t, []int{7, 8}, v1)
	assert.Equal(t, []int{1, 4, 9}, v2)
	assert.Equal(t, []int{4, 6}, v3)
	assert.Equal(t, []int{2, 4, 9}, v4)
	assert.Equal(t, []int{1, 2, 3, 4, 9}, v5)
	assert.Equal(t, []int{2, 5, 8, 9}, v6)
	assert.Equal(t, []int{8, 9}, v7)

}

func TestSolveLevel2(t *testing.T) {
	bs := `.......13
174..6..5
....5.6..
.....2.6.
4..9.5..1
.2.8.....
..3.7....
6..1..398
24.......`
	b := util.Parse(bs)

	sb, success := solve(b)

	expected := [][]int{
		{5, 6, 2, 7, 4, 8, 9, 1, 3},
		{1, 7, 4, 3, 9, 6, 2, 8, 5},
		{3, 9, 8, 2, 5, 1, 6, 4, 7},
		{7, 3, 5, 4, 1, 2, 8, 6, 9},
		{4, 8, 6, 9, 3, 5, 7, 2, 1},
		{9, 2, 1, 8, 6, 7, 5, 3, 4},
		{8, 1, 3, 6, 7, 9, 4, 5, 2},
		{6, 5, 7, 1, 2, 4, 3, 9, 8},
		{2, 4, 9, 5, 8, 3, 1, 7, 6}}
	assert.Equal(t, true, success)
	assert.Equal(t, expected, sb)

}

func TestSolveLevel1(t *testing.T) {
	bs := `.3.562.41
..48.....
6...497..
.86...32.
...6.7...
.75...19.
..921...8
.....65..
56.738.1.`
	b := util.Parse(bs)

	sb, success := solve(b)
	expected := [][]int{
		{9, 3, 7, 5, 6, 2, 8, 4, 1},
		{1, 5, 4, 8, 7, 3, 2, 6, 9},
		{6, 2, 8, 1, 4, 9, 7, 5, 3},
		{4, 8, 6, 9, 5, 1, 3, 2, 7},
		{3, 9, 1, 6, 2, 7, 4, 8, 5},
		{2, 7, 5, 3, 8, 4, 1, 9, 6},
		{7, 4, 9, 2, 1, 5, 6, 3, 8},
		{8, 1, 3, 4, 9, 6, 5, 7, 2},
		{5, 6, 2, 7, 3, 8, 9, 1, 4},
	}
	assert.Equal(t, true, success)
	assert.Equal(t, expected, sb)

}

func TestSolveLevel4(t *testing.T) {
	bs := `..14..7.9
.7.....1.
...2.....
35.87....
.8.9.5.2.
....46.38
.....9...
.6.....5.
5.3..16..`
	b := util.Parse(bs)

	sb, success := solve(b)

	expected := [][]int{
		{2, 3, 1, 4, 5, 8, 7, 6, 9},
		{8, 7, 5, 6, 9, 3, 2, 1, 4},
		{6, 4, 9, 2, 1, 7, 3, 8, 5},
		{3, 5, 4, 8, 7, 2, 1, 9, 6},
		{1, 8, 6, 9, 3, 5, 4, 2, 7},
		{9, 2, 7, 1, 4, 6, 5, 3, 8},
		{4, 1, 2, 5, 6, 9, 8, 7, 3},
		{7, 6, 8, 3, 2, 4, 9, 5, 1},
		{5, 9, 3, 7, 8, 1, 6, 4, 2},
	}
	assert.Equal(t, true, success)
	assert.Equal(t, expected, sb)

}

func TestSolveLevelMaster(t *testing.T) {
	bs := `.8.6...9.
...5..1..
..4.18.2.
7......4.
...1.....
.5..392..
.9..823..
..5.....9
.....6...`
	b := util.Parse(bs)

	sb, success := solve(b)

	expected := [][]int{
		{5, 8, 1, 6, 2, 3, 7, 9, 4},
		{3, 7, 2, 5, 9, 4, 1, 8, 6},
		{9, 6, 4, 7, 1, 8, 5, 2, 3},
		{7, 1, 3, 2, 6, 5, 9, 4, 8},
		{8, 2, 9, 1, 4, 7, 6, 3, 5},
		{4, 5, 6, 8, 3, 9, 2, 1, 7},
		{6, 9, 7, 4, 8, 2, 3, 5, 1},
		{2, 4, 5, 3, 7, 1, 8, 6, 9},
		{1, 3, 8, 9, 5, 6, 4, 7, 2},
	}
	assert.Equal(t, true, success)
	assert.Equal(t, expected, sb)

}
