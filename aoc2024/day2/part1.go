package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	f, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var result int

outer:
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if len(tokens) == 1 {
			result++
			continue
		}

		prev, _ := strconv.Atoi(tokens[0])
		curr, _ := strconv.Atoi(tokens[1])

		increasing := true
		if prev > curr {
			increasing = false
		}

		for i := 1; i < len(tokens); i++ {
			curr, _ := strconv.Atoi(tokens[i])
			diff := curr - prev

			if increasing && diff < 0 {
				continue outer
			} else if !increasing && diff > 0 {
				continue outer
			} else {
				if diff < 0 {
					diff = -diff
				}
				if diff < 1 || diff > 3 {
					continue outer
				}
			}
			prev = curr
		}
		result++
	}
	fmt.Println(result)
}
