package day19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part1() {
	f, err := os.Open("day19/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	scanner.Scan()
	var patterns []string = strings.Split(strings.TrimSuffix(scanner.Text(), "\n"), ", ")
	var patternsMap = make(map[string]bool)
	for _, p := range patterns {
		patternsMap[p] = true
	}
	scanner.Scan()

	var result int = 0
	for scanner.Scan() {
		design := strings.TrimSuffix(scanner.Text(), "\n")
		if isValidDesign(patternsMap, design) {
			result++
		}
	}
	fmt.Println("part1:", result)
}

func isValidDesign(patterns map[string]bool, design string) bool {
	dp := make([]bool, len(design)+1)
	for i := range dp {
		dp[i] = false
	}
	dp[len(dp)-1] = true

	for i := len(design) - 1; i >= 0; i-- {
		for j := i; j < len(design); j++ {
			if _, exists := patterns[design[i:j+1]]; exists {
				if dp[j+1] == true {
					dp[i] = true
				}
			}
		}
	}
	return dp[0] == true
}
