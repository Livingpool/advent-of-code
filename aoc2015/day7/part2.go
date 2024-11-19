package day7

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part2() {
	f, err := os.Open("day7/input.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	wires := make(map[string]*Node)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		tokens := strings.Split(line, " -> ")

		node := Node{
			filled: false,
			source: tokens[0],
		}
		wires[tokens[1]] = &node
	}

	result := decipher("a", wires)

	// reset all wires' filled boolean
	for key := range wires {
		wires[key].filled = false
	}

	wires["b"].val = result
	wires["b"].filled = true

	fmt.Println(decipher("a", wires))
}
