package day23

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Connection struct {
	city1 string
	city2 string
}

func Part1() {
	f, err := os.Open("day23/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var connections = make(map[Connection]bool)
	var computers = make(map[string]bool)

	for scanner.Scan() {
		tokens := strings.Split(strings.TrimSuffix(scanner.Text(), "\n"), "-")
		connections[Connection{tokens[0], tokens[1]}] = true
		connections[Connection{tokens[1], tokens[0]}] = true
		computers[tokens[0]] = true
		computers[tokens[1]] = true
	}

	var computerArray = make([]string, 0, len(computers))
	for computer := range computers {
		computerArray = append(computerArray, computer)
	}

	// loop: for all possible triplets of computers
	var result int = 0
	for i := 0; i < len(computerArray); i++ {
		for j := i + 1; j < len(computerArray); j++ {
			for k := j + 1; k < len(computerArray); k++ {
				// first check if they are all connected
				c1, c2, c3 := computerArray[i], computerArray[j], computerArray[k]
				if _, exists := connections[Connection{c1, c2}]; exists {
					if _, exists := connections[Connection{c2, c3}]; exists {
						if _, exists := connections[Connection{c1, c3}]; exists {
							// then check if at least one of them start with t
							if c1[0] == 't' || c2[0] == 't' || c3[0] == 't' {
								result++
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("part1:", result)
}
