package day8

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
	f, err := os.Open("day8/input.txt")
	if err != nil {
		panic(err)
	}

	var totalCode int = 0
	var totalEncoded int = 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		totalEncoded += 2
		for i := 0; i < len(line); i++ {
			switch rune(line[i]) {
			case '\\':
				totalEncoded += 2
			case '"':
				totalEncoded += 2
			default:
				totalEncoded++
			}
			totalCode++
		}
		// fmt.Printf("totalEncoded: %d, totalCode: %d\n", totalEncoded, totalCode)
	}

	fmt.Println(totalEncoded - totalCode)
}
