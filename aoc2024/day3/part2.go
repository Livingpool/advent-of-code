package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part2() {
	input, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	var str string = string(input)

	r, err := regexp.Compile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
	if err != nil {
		panic(err)
	}

	var result int = 0
	var enabled bool = true

	matches := r.FindAllString(str, -1)
	for _, m := range matches {
		if m[:3] == "don" {
			enabled = false
		} else if m[:2] == "do" {
			enabled = true
		} else if enabled {
			tokens := strings.Split(m[4:len(m)-1], ",")

			num1, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(err)
			}
			num2, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic(err)
			}
			result += num1 * num2
		}
	}
	fmt.Println(result)
}
