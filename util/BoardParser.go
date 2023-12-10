package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(f string) string {
	b, err := os.ReadFile(f) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	return str
}

func GetBoard(f string) [][]int {
	i := ReadFile(f)
	return Parse(i)
}

func Parse(input string) [][]int {
	result := [][]int{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), "")
		row := []int{}
		for _, p := range points {
			if p == "." {
				row = append(row, 0)
			} else {
				i, _ := strconv.Atoi(p)
				row = append(row, i)
			}
		}
		result = append(result, row)
	}
	return result
}

func Print(input [][]int) string {
	rows := []string{}
	rowH := "+-------+-------+-------+"
	for y, r := range input {
		if y%3 == 0 {
			rows = append(rows, rowH)
		}
		line := []string{}
		for x, c := range r {
			if x%3 == 0 {
				line = append(line, "|")
			}
			if c == 0 {
				line = append(line, ".")
			} else {
				line = append(line, strconv.Itoa(c))
			}
		}
		line = append(line, "|")
		rows = append(rows, strings.Join(line[:], " "))
	}
	rows = append(rows, rowH)
	return strings.Join(rows[:], "\n")
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
