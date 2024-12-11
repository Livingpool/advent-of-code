package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// hint: https://www.reddit.com/r/adventofcode/comments/1hbplc0/2024_day_11_how_far_can_you_go/

func Part2() {
	input, err := os.ReadFile("day11/input.txt")
	if err != nil {
		panic(err)
	}

	var transformations = make(map[string][]string)
	var frequency = make(map[string]int)

	tokens := strings.Split(strings.TrimSuffix(string(input), "\n"), " ")
	for _, t := range tokens {
		frequency[t]++
		if _, exists := transformations[t]; !exists {
			newStones, err := transform(t)
			transformations[t] = newStones
			if err != nil {
				panic(err)
			}
		}
	}

	// now for each round, we're simply updating the values.
	for range 75 {
		temp := make(map[string]int)
		for k, v := range frequency {
			if v > 0 {
				if stones, exists := transformations[k]; exists {
					for _, stone := range stones {
						temp[stone] += v
					}
				} else {
					newStones, err := transform(k)
					if err != nil {
						panic(err)
					}
					transformations[k] = newStones
					for _, stone := range newStones {
						temp[stone] += v
					}
				}
			}
		}
		frequency = temp
	}

	var result int = 0
	for _, v := range frequency {
		result += v
	}
	fmt.Println("part2:", result)
}

func transform(token string) ([]string, error) {
	if token == "0" {
		return []string{"1"}, nil
	} else if len(token)%2 == 0 {
		left := token[:len(token)/2]
		right := token[len(token)/2:]

		// remove leading zeroes
		rightNum, err := strconv.ParseUint(right, 10, 64)
		if err != nil {
			return nil, err
		}
		right = strconv.FormatUint(rightNum, 10)
		return []string{left, right}, nil
	} else {
		num, err := strconv.ParseUint(token, 10, 64)
		if err != nil {
			return nil, err
		}
		return []string{strconv.FormatUint(num*2024, 10)}, nil
	}

}
