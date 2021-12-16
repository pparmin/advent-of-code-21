package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	value  string
	marked bool
}

type Row struct {
	values     []Number
	horizontal int
}

type Board struct {
	rows []Row
	// holds a counter for each row; if one row reaches 5, we know it's full
	horizontal int
	// holds a counter for each position in each row; if one pos reaches 5
	// we know the column is full
	vertical map[int]int

	// only needed for part two
	ready bool
}

func newRow() *Row {
	r := Row{
		values:     make([]Number, 0),
		horizontal: 0,
	}
	return &r
}

func newBoard() *Board {
	b := Board{
		rows:       make([]Row, 0),
		horizontal: 0,
		vertical: map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
		},
	}
	return &b
}

func getBoards(input []string) []Board {
	b := newBoard()
	row := newRow()
	var (
		boards   []Board
		values   []string
		rowCount int = 0
	)

	for i, line := range input {
		if i == 0 {
			continue
		}
		if line == "" {
			continue
		}
		if rowCount > 4 {
			boards = append(boards, *b)
			rowCount = 0
			b = newBoard()
		}

		values = strings.Split(line, " ")
		for _, v := range values {
			if v == "" {
				continue
			} else {
				n := Number{
					value:  v,
					marked: false,
				}
				row.values = append(row.values, n)
			}
		}
		b.rows = append(b.rows, *row)
		row = newRow()
		rowCount++
	}
	return boards
}

func play(n string, boards []Board) (string, Board, bool) {
	for _, b := range boards {
		for k, r := range b.rows {
			for l, num := range r.values {
				if num.value == n {
					num.marked = true
					r.horizontal++
					b.vertical[l] += 1
					b.rows[k].values[l] = num
					b.rows[k].horizontal = r.horizontal
				}
			}
		}
		if isComplete(b) {
			return n, b, true
		}
	}
	return "", Board{}, false
}

func isComplete(b Board) bool {
	for _, v := range b.vertical {
		if v == 5 {
			return true
		}
	}

	for _, row := range b.rows {
		if row.horizontal == 5 {
			return true
		}
	}
	return false
}

func calcScore(b Board) int {
	sum := 0
	for _, row := range b.rows {
		for _, v := range row.values {
			if v.marked == false {
				val, err := strconv.Atoi(v.value)
				checkError(err)
				sum += val
			}
		}
	}
	return sum
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(name string) []string {
	file, err := os.Open(name)
	checkError(err)
	scanner := bufio.NewScanner(file)

	var text []string

	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}
	return text
}

func partOne() {
	input := ReadFile("input.txt")
	complete := false
	var final Board
	lastNum := ""

	numbers := strings.Split(input[0], ",")
	boards := getBoards(input)

	// TODO: Find a way to access each individual number in multi-
	// dimensional array
	for _, n := range numbers {
		lastNum, final, complete = play(n, boards)
		if complete {
			fmt.Println("FINAL BOARD:", final)
			result := calcScore(final)
			lNum, err := strconv.Atoi(lastNum)
			checkError(err)
			fmt.Println("FINAL RESULT: ", result*lNum)
			return
		}
	}
}

func main() {
	partOne()
}
