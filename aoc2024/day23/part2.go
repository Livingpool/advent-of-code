package day23

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// hint: https://www.reddit.com/r/adventofcode/comments/1hkgj5b/comment/m3egbtj/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
func Part2() {
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

	// we initialize a list of singleton sets; each computer has its own network
	var networks = make(map[string]map[string]bool)
	for c := range computers {
		networks[c] = make(map[string]bool)
		networks[c][c] = true
	}

	// this greedy approach works on the input for some reason...
	for i := range networks {
	outer:
		for c := range computers {
			for d := range networks[i] {
				if _, exists := connections[Connection{c, d}]; !exists {
					continue outer
				}
			}
			networks[i][c] = true
		}
	}

	var count int
	var node string
	for i := range networks {
		if len(networks[i]) > count {
			count = len(networks[i])
			node = i
		}
	}

	var result []string
	for k := range networks[node] {
		result = append(result, k)
	}
	slices.Sort(result)
	fmt.Println(strings.Join(result, ","))
}
