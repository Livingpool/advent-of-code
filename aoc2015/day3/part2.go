package day3

import (
	"fmt"
	"os"
)

func Part2() {
	input, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	santa := Pair{First: 0, Second: 0}
	robo := Pair{First: 0, Second: 0}
	m := make(map[Pair]bool)
	m[santa] = true

	result := 1

	for i, c := range string(input) {
		switch c {
		case '^':
			if i%2 == 0 {
				santa.Second++
			} else {
				robo.Second++
			}
		case '>':
			if i%2 == 0 {
				santa.First++
			} else {
				robo.First++
			}
		case 'v':
			if i%2 == 0 {
				santa.Second--
			} else {
				robo.Second--
			}
		case '<':
			if i%2 == 0 {
				santa.First--
			} else {
				robo.First--
			}
		}

		if i%2 == 0 {
			if _, exists := m[santa]; !exists {
				m[santa] = true
				result++
			}
		} else {
			if _, exists := m[robo]; !exists {
				m[robo] = true
				result++
			}
		}
	}

	fmt.Println(result)
}
