package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1() {
	input, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	var str string = string(input)

	r, err := regexp.Compile(`mul\(\d+,\d+\)`)
	if err != nil {
		panic(err)
	}

	var result int = 0

	matches := r.FindAllString(str, -1)
	for _, m := range matches {
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
	fmt.Println(result)
}
