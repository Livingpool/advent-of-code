package day8

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() {
	f, err := os.Open("day8/input.txt")
	if err != nil {
		panic(err)
	}

	var totalCode int = 0
	var totalChar int = 0

	var backslashInEffect bool = false

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		for i := 0; i < len(line); i++ {
			if i != 0 && i != len(line)-1 {
				switch rune(line[i]) {
				case '\\':
					if backslashInEffect {
						totalChar++
						backslashInEffect = false
					} else {
						backslashInEffect = true
					}
				case 'x':
					if backslashInEffect {
						i += 2
						totalCode += 3
						totalChar++
						backslashInEffect = false
						continue
					} else {
						totalChar++
					}
				case '"':
					totalChar++
					backslashInEffect = false
				default:
					totalChar++
				}
			}
			totalCode++
		}
		// fmt.Printf("totalCode: %d, totalChar: %d\n", totalCode, totalChar)
	}

	fmt.Println(totalCode - totalChar)
}
