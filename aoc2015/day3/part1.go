package day3

import (
	"fmt"
	"os"
)

type Pair struct {
	First  int
	Second int
}

func Part1() {
	input, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	pos := Pair{First: 0, Second: 0}
	m := make(map[Pair]bool)
	m[pos] = true

	result := 1

	for _, c := range string(input) {
		switch c {
		case '^':
			pos.Second++
		case '>':
			pos.First++
		case 'v':
			pos.Second--
		case '<':
			pos.First--
		}

		if _, exists := m[pos]; !exists {
			m[pos] = true
			result++
		}
	}

	fmt.Println(result)
}
