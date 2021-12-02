package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func makeSum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func partOne() {
	file, err := os.Open("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(file)

	var (
		increases  int = 0
		prevNumber int = 0
	)

	for scanner.Scan() {
		curNumber, err := strconv.Atoi(scanner.Text())
		checkError(err)
		if prevNumber < curNumber && prevNumber != 0 {
			increases++
		}
		prevNumber = curNumber
	}

	checkError(scanner.Err())
	fmt.Println("Number of increases: ", increases)
}

func partTwo() {
	file, err := os.Open("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(file)

	var (
		values    []int
		sums      []int
		start     int = 0
		end       int = 3
		prevSum   int = 0
		curSum    int = 0
		increases int = 0
	)

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		checkError(err)
		values = append(values, number)
	}

	for i := 0; i < len(values); i++ {
		if end > len(values) {
			break
		}

		curSum = makeSum(values[start:end])
		sums = append(sums, curSum)

		if prevSum < curSum && prevSum != 0 {
			increases++
			prevSum = curSum
		}

		prevSum = curSum
		end++
		start++
	}

	checkError(scanner.Err())
	fmt.Println("No of increases: ", increases)
}

func main() {
	partOne()
	partTwo()
}
