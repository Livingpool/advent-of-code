package day19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part2() {
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
		result += countDesigns(patternsMap, design)
	}
	fmt.Println("part2:", result)
}

func countDesigns(patterns map[string]bool, design string) int {
	dp := make([]int, len(design)+1)
	dp[len(dp)-1] = 1

	for i := len(design) - 1; i >= 0; i-- {
		for j := i; j < len(design); j++ {
			if _, exists := patterns[design[i:j+1]]; exists {
				dp[i] += dp[j+1]
			}
		}
	}
	return dp[0]
}
