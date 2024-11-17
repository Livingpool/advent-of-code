package day1

import (
	"fmt"
	"os"
)

func Part1() {
	input, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}

	result := 0
	for _, c := range string(input) {
		if c == '(' {
			result++
		} else if c == ')' {
			result--
		}
	}

	fmt.Println(result)
}
