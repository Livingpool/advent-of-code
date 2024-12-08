package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part2() {
	f, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}

	var rules = make(map[Key]bool)

	// construct page ordering rules
	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		tokens := strings.Split(line, "|")
		first, second := tokens[0], tokens[1]
		num1, _ := strconv.Atoi(first)
		num2, _ := strconv.Atoi(second)

		rules[Key{a: num1, b: num2}] = true
	}

	var result int = 0

	compare := func(a, b int) int {
		if _, exists := rules[Key{a: a, b: b}]; exists {
			return -1
		} else {
			return 0
		}
	}

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		update := make([]int, len(tokens))
		for i, t := range tokens {
			val, _ := strconv.Atoi(t)
			update[i] = val
		}

		sorted := make([]int, len(update))
		copy(sorted, update)
		slices.SortStableFunc(sorted, compare)

		incorrect := false
		for i := range update {
			if update[i] != sorted[i] {
				incorrect = true
				break
			}
		}
		if incorrect {
			result += getMid(sorted)
		}
	}

	fmt.Println("total:", result)
}
