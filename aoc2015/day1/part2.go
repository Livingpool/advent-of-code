package day1

import (
	"fmt"
	"os"
)

func Part2() {
	input, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}

	result := 0
	for pos, c := range string(input) {
		if c == '(' {
			result++
		} else if c == ')' {
			result--
		}

		if result < 0 {
			fmt.Println(pos + 1)
			return
		}
	}

	fmt.Println(result)
}
