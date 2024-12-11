package day11

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	input, err := os.ReadFile("day11/input.txt")
	if err != nil {
		panic(err)
	}

	tokens := strings.Split(strings.TrimSuffix(string(input), "\n"), " ")
	for i := 0; i < 25; i++ {
		temp := make([]string, len(tokens))
		copy(temp, tokens)

		k := 0

		for j := 0; j < len(tokens); j++ {
			if len(tokens[j]) == 1 && tokens[j] == "0" {
				temp[k] = "1"
				k++
			} else if len(tokens[j])%2 == 0 {
				left := tokens[j][:len(tokens[j])/2]
				right := tokens[j][len(tokens[j])/2:]

				// remove leading zeroes
				rightNum, err := strconv.ParseUint(right, 10, 64)
				if err != nil {
					panic(err)
				}
				right = strconv.FormatUint(rightNum, 10)

				temp[k] = left
				temp = slices.Insert(temp, k+1, right)
				k += 2
			} else {
				num, err := strconv.ParseUint(tokens[j], 10, 64)
				if err != nil {
					panic(err)
				}

				newNum := num * 2024
				temp[k] = strconv.FormatUint(newNum, 10)
				k++
			}
		}
		tokens = temp
	}
	fmt.Println("part1:", len(tokens))
}
