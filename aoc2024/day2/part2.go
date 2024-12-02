package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	f, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var result int

outer:
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if len(tokens) == 1 {
			result++
			continue
		}

		prev, _ := strconv.Atoi(tokens[0])
		curr, _ := strconv.Atoi(tokens[1])

		increasing := true
		if prev > curr {
			increasing = false
		}

		cand1, cand2, cand3 := -1, -1, -1

	inner:
		for i := 1; i < len(tokens); i++ {
			curr, _ := strconv.Atoi(tokens[i])
			diff := curr - prev

			if increasing && diff < 0 {
				cand1, cand2, cand3 = i-1, i, i-2
				break inner
			} else if !increasing && diff > 0 {
				cand1, cand2, cand3 = i-1, i, i-2
				break inner
			} else {
				if diff < 0 {
					diff = -diff
				}
				if diff < 1 || diff > 3 {
					cand1, cand2 = i-1, i
					break inner
				}
			}
			prev = curr
		}

		temp := make([]string, len(tokens))
		temp2 := make([]string, len(tokens))
		copy(temp, tokens)
		copy(temp2, tokens)

		if cand1 >= 0 && !isSafe(temp, cand1) && !isSafe(temp2, cand2) {
			if cand3 >= 0 {
				if !isSafe(tokens, cand3) {
					continue outer
				}
			} else {
				continue outer
			}
		}
		result++
	}
	fmt.Println(result)
}

func isSafe(tokens []string, ignore int) bool {
	tokens = append(tokens[:ignore], tokens[ignore+1:]...)

	prev, _ := strconv.Atoi(tokens[0])
	curr, _ := strconv.Atoi(tokens[1])

	increasing := true
	if prev > curr {
		increasing = false
	}

	for i := 1; i < len(tokens); i++ {
		curr, _ := strconv.Atoi(tokens[i])
		diff := curr - prev

		if increasing && diff < 0 {
			return false
		} else if !increasing && diff > 0 {
			return false
		} else {
			if diff < 0 {
				diff = -diff
			}
			if diff < 1 || diff > 3 {
				return false
			}
		}
		prev = curr
	}
	return true
}
