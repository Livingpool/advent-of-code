package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var left []int
	var right = make(map[int]int)

	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		leftVal, _ := strconv.Atoi(tokens[0])
		rightVal, _ := strconv.Atoi(tokens[1])

		left = append(left, leftVal)
		right[rightVal]++
	}

	var result int = 0
	for i := 0; i < len(left); i++ {
		if val, exists := right[left[i]]; exists {
			result += left[i] * val
		}
	}
	fmt.Println(result)
}
