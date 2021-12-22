package day5

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y string
}

type Line struct {
	start Point
	end   Point
}

func NewPoint() *Point {
	p := Point{X: "", Y: ""}
	return &p
}

func NewLine() *Line {
	l := Line{start: Point{}, end: Point{}}
	return &l
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(name string) []string {
	file, err := os.Open(name)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}
	return text
}

func parseNumber(line string) []string {
	var number []rune
	var result []string

	for i, char := range line {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			number = append(number, char)
			// handle end of line
			if i == len(line)-1 {
				result = append(result, string(number))
			}
		case ',', '>':
			result = append(result, string(number))
			number = []rune{}
		case ' ', '-':
			continue
		}
	}
	return result
}

// parse input to generate the field of hydrothermal vents
func GetLines(input []string) []Line {
	var lines []Line
	for _, line := range input {
		fmt.Println("LINE: ", line)
		l := NewLine()
		p := NewPoint()

		numbers := parseNumber(line)
		fmt.Println("Numbers: ", numbers)
		p.X = numbers[0]
		p.Y = numbers[1]
		fmt.Println("CURRENT POINT: ", p)
		l.start = *p
		fmt.Println("CURRENT LINE: ", l)

		p.X = numbers[2]
		p.Y = numbers[3]
		fmt.Println("CURRENT POINT: ", p)
		l.end = *p
		fmt.Println("CURRENT LINE: ", l)
		lines = append(lines, *l)
	}
	return lines
}

func CalcOverlap(lines []Line) int {
	return 0
}
