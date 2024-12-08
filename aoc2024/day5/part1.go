package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Key struct {
	a int
	b int
}

// a rule X | Y means that both X and Y have to be produced as part of the update
// thus, i don't need to consider transitive property
// so simple sorting works...

func Part1() {
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

outer:
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

		for i := range update {
			if update[i] != sorted[i] {
				continue outer
			}
		}
		result += getMid(update)
	}

	fmt.Println("total:", result)
}

func getMid(items []int) int {
	return items[len(items)/2]
}
