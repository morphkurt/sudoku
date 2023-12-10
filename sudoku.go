package main

// Hello returns a greeting for the named person.

func solve(board [][]int) (b [][]int, solved bool) {
	r, c := getNextEmpty(board)
	if r == -1 && c == -1 {
		if IsSolved(board) {
			return board, true
		}
		return board, false
	}
	vals := GetPossibleValues(board, c, r)
	for _, p := range vals {
		board[r][c] = p
		b, s := solve(board)
		if s {
			return b, s
		} else {
			board[r][c] = 0
		}
	}

	return board, false
}

func getNextEmpty(board [][]int) (r, c int) {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			v := board[y][x]
			if v == 0 {
				return y, x
			}
		}
	}
	return -1, -1
}

func IsSolved(board [][]int) bool {
	cr := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	cc := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	cb := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			v := board[x][y]
			if v == 0 {
				return false
			}
			cr[y][v-1] += 1
			cc[x][v-1] += 1
			cb[(((y / 3) * 3) + x/3)][v-1] += 1
		}
	}
	if IsOnes(cc) && IsOnes(cr) {
		return true
	}
	return false
}

func GetPossibleValues(board [][]int, c int, r int) []int {
	cr := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	cc := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	cb := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			v := board[y][x]
			if v > 0 {
				cr[y][v-1] += 1
				cc[x][v-1] += 1
				cb[(((y / 3) * 3) + x/3)][v-1] += 1
			}
		}
	}
	pcr := GetValues(cr[r], 0)
	pcc := GetValues(cc[c], 0)
	pcb := GetValues(cb[(((r/3)*3)+c/3)], 0)

	return FindIntersection(pcr, pcc, pcb)

}

func IsOnes(c [][]int) bool {
	allOnes := true
	for y := 0; y < len(c); y++ {
		for x := 0; x < len(c[y]); x++ {
			if c[x][y] != 1 {
				return false
			}
		}
	}
	return allOnes
}

func GetValues(c []int, e int) []int {
	r := []int{}
	for i, v := range c {
		if v == e {
			r = append(r, i+1)
		}
	}
	return r
}

func FindIntersection(c1 []int, c2 []int, c3 []int) []int {
	v := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, i := range c1 {
		v[i-1] += 1
	}
	for _, i := range c2 {
		v[i-1] += 1
	}
	for _, i := range c3 {
		v[i-1] += 1
	}
	return GetValues(v, 3)
}
