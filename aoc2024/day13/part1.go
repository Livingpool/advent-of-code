package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type button struct {
	x int
	y int
}

func Part1() {
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
		scanner.Scan() // discard newline

		t := calcFewestTokens(btnA, btnB, prize)
		if t > 0 {
			result += t
		}
	}
	fmt.Println("part1:", result)
}

func readInput(line string) (*button, error) {
	tokens := strings.Split(strings.TrimSuffix(line, "\n"), ": ")
	positions := strings.Split(tokens[1], ", ")

	x, err := strconv.Atoi(positions[0][2:])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(positions[1][2:])
	if err != nil {
		return nil, err
	}

	return &button{x: x, y: y}, nil
}

func calcFewestTokens(btnA, btnB, prize *button) int {
	a, b, k := btnA.x+btnA.y, btnB.x+btnB.y, prize.x+prize.y
	maxA := min(100, k/a)
	maxB := min(100, k/b)
	minA := max(0, (k-b*maxB)/a)

	for i := minA; i <= maxA; i++ {
		if (k-a*i)%b == 0 {
			// note: combining two equations into one actually loses information!
			j := (k - a*i) / b
			if i*btnA.x+j*btnB.x == prize.x && i*btnA.y+j*btnB.y == prize.y {
				return 3*i + j
			}
		}
	}
	return -1
}
