package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var left, right []int

	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		leftVal, _ := strconv.Atoi(tokens[0])
		rightVal, _ := strconv.Atoi(tokens[1])

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	slices.Sort(left)
	slices.Sort(right)

	var result int = 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff >= 0 {
			result += diff
		} else {
			result += -diff
		}
	}
	fmt.Println(result)
}
