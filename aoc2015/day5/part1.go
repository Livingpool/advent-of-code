package day5

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() {
	f, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}

	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	naughtyStrings := map[string]bool{
		"ab": true,
		"cd": true,
		"pq": true,
		"xy": true,
	}

	result := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		cond1 := 0
		cond2, cond3 := false, true

		for i := 0; i < len(line); i++ {
			if _, exists := vowels[line[i]]; exists {
				cond1++
			}
			if i < len(line)-1 {
				if line[i] == rune(line[i+1]) {
					cond2 = true
				}
				if _, exists := naughtyStrings[string(line[i:i+2])]; exists {
					cond3 = false
				}
			}
		}
		if cond1 >= 3 && cond2 && cond3 {
			result++
		}
	}
	fmt.Println(result)
}
