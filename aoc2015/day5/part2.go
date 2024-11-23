package day5

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
	f, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}

	var result int = 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		cond1, cond2 := false, false
		twoLetters := make(map[string]int) // 2-letter string, starting jposition

		for i := 0; i < len(line); i++ {
			if cond1 && cond2 {
				break
			} else {
				if !cond1 && i < len(line)-1 {
					if pos, exists := twoLetters[line[i:i+2]]; exists {
						if pos+1 < i {
							cond1 = true
						}
					} else {
						twoLetters[line[i:i+2]] = i
					}
				}

				if !cond2 && i > 0 && i < len(line)-1 {
					if line[i-1] == line[i+1] {
						cond2 = true
					}
				}
			}
		}

		if cond1 && cond2 {
			result++
		}
	}
	fmt.Println(result)
}
