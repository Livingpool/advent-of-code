package day13

import (
	"bufio"
	"fmt"
	"os"
)

// hint: https://www.reddit.com/r/adventofcode/comments/1hd5b6o/2024_day_13_in_the_end_math_reigns_supreme/
func Part2() {
	f, err := os.Open("day13/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var result int = 0
	for scanner.Scan() {
		btnA, _ := readInput(scanner.Text())
		scanner.Scan()
		btnB, _ := readInput(scanner.Text())
		scanner.Scan()
		prize, _ := readInput(scanner.Text())
		prize.x += 10000000000000
		prize.y += 10000000000000
		scanner.Scan() // discard newline

		t := linearAlgebra(btnA, btnB, prize)
		if t > 0 {
			result += t
		}
	}
	fmt.Println("part2:", result)
}

/*
|A.x B.x| |A| = |P.x|
|A.y B.y| |B|   |P.y|

Two equations, two variables => uniquely solve for A, B!
*/
func linearAlgebra(btnA, btnB, prize *button) int {
	determinant := btnA.x*btnB.y - btnA.y*btnB.x
	if determinant == 0 { // no inverse of the matrix
		return -1
	}

	// for 2d matrix
	rowA := btnB.y*prize.x - btnB.x*prize.y
	rowB := (-btnA.y * prize.x) + btnA.x*prize.y

	if rowA%determinant != 0 || rowB%determinant != 0 {
		return -1
	}

	A := rowA / determinant
	B := rowB / determinant

	if A < 0 || B < 0 {
		return -1
	}

	return 3*A + B
}
