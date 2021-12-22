package day5_test

import (
	d5 "aoc/day5"
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	t.Parallel()
	t.Run("test input part one", func(t *testing.T) {
		input := d5.ReadFile("test.txt")
		got := d5.GetLines(input)
		fmt.Println("GOT: ", got)
	})

	t.Run("real input part one", func(t *testing.T) {
		input := d5.ReadFile("input.txt")
		got := d5.GetLines(input)
		fmt.Println("GOT: ", got)
	})
}
