package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// It's important to note that the input is a list of relationships, not a series of steps to follow.
// This means starting with wire `a` and tracing the inputs to `a` to their inputs, and so on.

// So i think i will build a tree, and then recursively get the value of node 'a'!

type Node struct {
	val    uint16
	filled bool
	source string
}

func Part1() {
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

	fmt.Println(decipher("a", wires))
}

func decipher(s string, wires map[string]*Node) uint16 {
	// first check is s is a single int
	if len(strings.Split(s, " ")) == 1 && unicode.IsDigit(rune(s[0])) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return uint16(num)
	}

	// if not, continue
	curr, exists := wires[s]
	if !exists {
		fmt.Println("error: ", s, " not found in map")
		return 0
	}

	var result uint16

	if curr.filled {
		return curr.val
	} else {
		tokens := strings.Split(curr.source, " ")

		if len(tokens) == 1 { // 123 -> x, y -> x
			if unicode.IsDigit(rune(s[0])) {
				num, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				result = uint16(num)
			} else {
				result = decipher(tokens[0], wires)
			}
		} else if len(tokens) == 2 { // NOT x
			num := decipher(tokens[1], wires)
			return ^num
		} else { // AND, OR, LSHIFT, RSHIFT
			switch tokens[1] {
			case "AND":
				result = decipher(tokens[0], wires) & decipher(tokens[2], wires)
			case "OR":
				result = decipher(tokens[0], wires) | decipher(tokens[2], wires)
			case "LSHIFT":
				num, err := strconv.Atoi(tokens[2])
				if err != nil {
					panic(err)
				}
				result = decipher(tokens[0], wires) << num
			case "RSHIFT":
				num, err := strconv.Atoi(tokens[2])
				if err != nil {
					panic(err)
				}
				result = decipher(tokens[0], wires) >> num
			default:
				fmt.Println("error: uncatched operator")
				result = 0
			}
		}
	}

	curr.filled = true
	curr.val = result
	return result
}
